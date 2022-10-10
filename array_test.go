// Copyright 2014 Unknwon
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package featws

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestArray_NewItem(t *testing.T) {
	t.Run("create a new item", func(t *testing.T) {
		f := Empty()
		require.NotNil(t, f)

		a := f.Array("[array]")

		s, err := a.NewItem()
		require.NoError(t, err)
		require.NotNil(t, s)
		assert.Equal(t, "[array]"+strconv.Itoa(a.Length()-1), s.Name())

		t.Run("with empty string", func(t *testing.T) {
			_, err := f.ArrayByName("")
			require.Error(t, err)
		})
	})
}

// func TestArray_GetKey(t *testing.T) {
// 	t.Run("get a key", func(t *testing.T) {
// 		f := Empty()
// 		require.NotNil(t, f)

// 		k, err := f.Array("").NewKey("NAME", "ini")
// 		require.NoError(t, err)
// 		require.NotNil(t, k)

// 		k, err = f.Array("").GetKey("NAME")
// 		require.NoError(t, err)
// 		require.NotNil(t, k)
// 		assert.Equal(t, "NAME", k.Name())
// 		assert.Equal(t, "ini", k.Value())

// 		t.Run("key not exists", func(t *testing.T) {
// 			_, err := f.Array("").GetKey("404")
// 			require.Error(t, err)
// 		})

// 		t.Run("key exists in parent section", func(t *testing.T) {
// 			k, err := f.Array("parent").NewKey("AGE", "18")
// 			require.NoError(t, err)
// 			require.NotNil(t, k)

// 			k, err = f.Array("parent.child.son").GetKey("AGE")
// 			require.NoError(t, err)
// 			require.NotNil(t, k)
// 			assert.Equal(t, "18", k.Value())
// 		})
// 	})
// }

// func TestArray_HasKey(t *testing.T) {
// 	t.Run("check if a key exists", func(t *testing.T) {
// 		f := Empty()
// 		require.NotNil(t, f)

// 		k, err := f.Array("").NewKey("NAME", "ini")
// 		require.NoError(t, err)
// 		require.NotNil(t, k)

// 		assert.True(t, f.Array("").HasKey("NAME"))
// 		assert.True(t, f.Array("").HasKey("NAME"))
// 		assert.False(t, f.Array("").HasKey("404"))
// 		assert.False(t, f.Array("").HasKey("404"))
// 	})
// }

// func TestArray_HasValue(t *testing.T) {
// 	t.Run("check if contains a value in any key", func(t *testing.T) {
// 		f := Empty()
// 		require.NotNil(t, f)

// 		k, err := f.Array("").NewKey("NAME", "ini")
// 		require.NoError(t, err)
// 		require.NotNil(t, k)

// 		assert.True(t, f.Array("").HasValue("ini"))
// 		assert.False(t, f.Array("").HasValue("404"))
// 	})
// }

// func TestArray_Key(t *testing.T) {
// 	t.Run("get a key", func(t *testing.T) {
// 		f := Empty()
// 		require.NotNil(t, f)

// 		k, err := f.Array("").NewKey("NAME", "ini")
// 		require.NoError(t, err)
// 		require.NotNil(t, k)

// 		k = f.Array("").Key("NAME")
// 		require.NotNil(t, k)
// 		assert.Equal(t, "NAME", k.Name())
// 		assert.Equal(t, "ini", k.Value())

// 		t.Run("key not exists", func(t *testing.T) {
// 			k := f.Array("").Key("404")
// 			require.NotNil(t, k)
// 			assert.Equal(t, "404", k.Name())
// 		})

// 		t.Run("key exists in parent section", func(t *testing.T) {
// 			k, err := f.Array("parent").NewKey("AGE", "18")
// 			require.NoError(t, err)
// 			require.NotNil(t, k)

// 			k = f.Array("parent.child.son").Key("AGE")
// 			require.NotNil(t, k)
// 			assert.Equal(t, "18", k.Value())
// 		})
// 	})
// }

// func TestArray_Keys(t *testing.T) {
// 	t.Run("get all keys in a section", func(t *testing.T) {
// 		f := Empty()
// 		require.NotNil(t, f)

// 		k, err := f.Array("").NewKey("NAME", "ini")
// 		require.NoError(t, err)
// 		require.NotNil(t, k)
// 		k, err = f.Array("").NewKey("VERSION", "v1")
// 		require.NoError(t, err)
// 		require.NotNil(t, k)
// 		k, err = f.Array("").NewKey("IMPORT_PATH", "gopkg.in/ini.v1")
// 		require.NoError(t, err)
// 		require.NotNil(t, k)

// 		keys := f.Array("").Keys()
// 		names := []string{"NAME", "VERSION", "IMPORT_PATH"}
// 		assert.Equal(t, len(names), len(keys))
// 		for i, name := range names {
// 			assert.Equal(t, name, keys[i].Name())
// 		}
// 	})
// }

// func TestArray_ParentKeys(t *testing.T) {
// 	t.Run("get all keys of parent sections", func(t *testing.T) {
// 		f := Empty()
// 		require.NotNil(t, f)

// 		k, err := f.Array("package").NewKey("NAME", "ini")
// 		require.NoError(t, err)
// 		require.NotNil(t, k)
// 		k, err = f.Array("package").NewKey("VERSION", "v1")
// 		require.NoError(t, err)
// 		require.NotNil(t, k)
// 		k, err = f.Array("package").NewKey("IMPORT_PATH", "gopkg.in/ini.v1")
// 		require.NoError(t, err)
// 		require.NotNil(t, k)

// 		keys := f.Array("package.sub.sub2").ParentKeys()
// 		names := []string{"NAME", "VERSION", "IMPORT_PATH"}
// 		assert.Equal(t, len(names), len(keys))
// 		for i, name := range names {
// 			assert.Equal(t, name, keys[i].Name())
// 		}
// 	})
// }

// func TestArray_KeyStrings(t *testing.T) {
// 	t.Run("get all key names in a section", func(t *testing.T) {
// 		f := Empty()
// 		require.NotNil(t, f)

// 		k, err := f.Array("").NewKey("NAME", "ini")
// 		require.NoError(t, err)
// 		require.NotNil(t, k)
// 		k, err = f.Array("").NewKey("VERSION", "v1")
// 		require.NoError(t, err)
// 		require.NotNil(t, k)
// 		k, err = f.Array("").NewKey("IMPORT_PATH", "gopkg.in/ini.v1")
// 		require.NoError(t, err)
// 		require.NotNil(t, k)

// 		assert.Equal(t, []string{"NAME", "VERSION", "IMPORT_PATH"}, f.Array("").KeyStrings())
// 	})
// }

func TestArray_KeyHash(t *testing.T) {
	t.Run("get clone of key hash", func(t *testing.T) {
		f, err := Load([]byte(`
key = two
[[array]]
name = app1
file = a.log
[[array]]
name = app2
file = a.log
`), []byte(`
[[array]]
name = app3
file = a.log
`))
		require.NoError(t, err)
		require.NotNil(t, f)

		assert.Equal(t, "two", f.Section("").Key("key").String())

		assert.Equal(t, 3, f.Array("[array]").Length())
		// relations := []map[string]string{{
		// 	"name": "app2",
		// 	"file": "b.log",
		// }}
		// for k, v := range relations {
		// 	assert.Equal(t, relation[k], v)
		// }
	})
}

// func TestArray_DeleteKey(t *testing.T) {
// 	t.Run("delete a key", func(t *testing.T) {
// 		f := Empty()
// 		require.NotNil(t, f)

// 		k, err := f.Array("").NewKey("NAME", "ini")
// 		require.NoError(t, err)
// 		require.NotNil(t, k)

// 		assert.True(t, f.Array("").HasKey("NAME"))
// 		f.Array("").DeleteKey("NAME")
// 		assert.False(t, f.Array("").HasKey("NAME"))
// 	})
// }
