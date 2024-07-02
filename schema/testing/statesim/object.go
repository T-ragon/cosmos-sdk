package statesim

import (
	"fmt"

	"github.com/tidwall/btree"
	"pgregory.net/rapid"

	"cosmossdk.io/schema"
	schematesting "cosmossdk.io/schema/testing"
)

// ObjectCollection is a collection of objects of a specific type for testing purposes.
type ObjectCollection struct {
	options           Options
	objectType        schema.ObjectType
	objects           *btree.Map[string, schema.ObjectUpdate]
	updateGen         *rapid.Generator[schema.ObjectUpdate]
	valueFieldIndices map[string]int
}

// NewObjectCollection creates a new ObjectCollection for the given object type.
func NewObjectCollection(objectType schema.ObjectType, options Options) *ObjectCollection {
	objects := &btree.Map[string, schema.ObjectUpdate]{}
	updateGen := schematesting.ObjectUpdateGen(objectType, objects)
	valueFieldIndices := make(map[string]int, len(objectType.ValueFields))
	for i, field := range objectType.ValueFields {
		valueFieldIndices[field.Name] = i
	}

	return &ObjectCollection{
		options:           options,
		objectType:        objectType,
		objects:           objects,
		updateGen:         updateGen,
		valueFieldIndices: valueFieldIndices,
	}
}

// ApplyUpdate applies the given object update to the collection.
func (o *ObjectCollection) ApplyUpdate(update schema.ObjectUpdate) error {
	if update.TypeName != o.objectType.Name {
		return fmt.Errorf("update type name %q does not match object type name %q", update.TypeName, o.objectType.Name)
	}

	err := o.objectType.ValidateObjectUpdate(update)
	if err != nil {
		return err
	}

	keyStr := fmt.Sprintf("%v", update.Key)
	cur, exists := o.objects.Get(keyStr)
	if update.Delete {
		if o.objectType.RetainDeletions && o.options.CanRetainDeletions {
			if !exists {
				return fmt.Errorf("object not found for deletion: %v", update.Key)
			}

			cur.Delete = true
			o.objects.Set(keyStr, cur)
		} else {
			o.objects.Delete(keyStr)
		}
	} else {
		// merge value updates only if we have more than one value field
		if valueUpdates, ok := update.Value.(schema.ValueUpdates); ok &&
			len(o.objectType.ValueFields) > 1 {
			var values []interface{}
			if exists {
				values = cur.Value.([]interface{})
			} else {
				values = make([]interface{}, len(o.objectType.ValueFields))
			}

			err = valueUpdates.Iterate(func(fieldName string, value interface{}) bool {
				fieldIndex, ok := o.valueFieldIndices[fieldName]
				if !ok {
					panic(fmt.Sprintf("field %q not found in object type %q", fieldName, o.objectType.Name))
				}

				values[fieldIndex] = value
				return true
			})
			if err != nil {
				return err
			}

			update.Value = values
		}

		o.objects.Set(keyStr, update)
	}

	return nil
}

// UpdateGen returns a generator for random object updates against the collection. This generator
// is stateful and returns a certain number of updates and deletes to existing objects.
func (o *ObjectCollection) UpdateGen() *rapid.Generator[schema.ObjectUpdate] {
	return o.updateGen
}

// ScanState scans the state of the collection by calling the given function for each object update.
func (o *ObjectCollection) ScanState(f func(schema.ObjectUpdate) error) error {
	var err error
	o.objects.Scan(func(_ string, v schema.ObjectUpdate) bool {
		err = f(v)
		return err == nil
	})
	return err
}

// GetObject returns the object with the given key from the collection represented as an ObjectUpdate
// itself. Deletions that are retained are returned as ObjectUpdate's with delete set to true.
func (o *ObjectCollection) GetObject(key any) (update schema.ObjectUpdate, found bool) {
	return o.objects.Get(fmt.Sprintf("%v", key))
}

// ObjectType returns the object type of the collection.
func (o *ObjectCollection) ObjectType() schema.ObjectType {
	return o.objectType
}

// Len returns the number of objects in the collection.
func (o *ObjectCollection) Len() int {
	return o.objects.Len()
}
