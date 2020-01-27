// Copyright © 2020 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import Marshaler from '.'

describe('Marshaler', () => {
  describe('pathsFromPatch', () => {
    it('should handle undefined patch argument and produce empty paths', () => {
      const patch = undefined

      const paths = Marshaler.pathsFromPatch(patch)

      expect(paths).toBeDefined()
      expect(paths).toHaveLength(0)
    })

    it('should produce empty paths from empty patch', () => {
      const patch = {}

      const paths = Marshaler.pathsFromPatch(patch)

      expect(paths).toBeDefined()
      expect(paths).toHaveLength(0)
    })

    it('should produce paths from flat patch', () => {
      const fieldName1 = 'patchField1'
      const fieldName2 = 'patchField2'

      const patch = { [fieldName1]: 'test', [fieldName2]: 'test' }

      const paths = Marshaler.pathsFromPatch(patch)

      expect(paths).toBeDefined()
      expect(paths).toHaveLength(2)
    })

    it('should produce paths from nested patch', () => {
      const fieldName1 = 'patchField1'
      const fieldName2 = 'patchField2'
      const fieldName3 = 'nestedPatchField3'

      const patch = {
        [fieldName1]: {
          [fieldName3]: 'test',
        },
        [fieldName2]: 'test',
      }

      const paths = Marshaler.pathsFromPatch(patch)

      expect(paths).toBeDefined()
      expect(paths).toHaveLength(2)
      expect(paths).toEqual(expect.arrayContaining([[fieldName1, fieldName3], [fieldName2]]))
    })

    it('should produce paths from patch containing arrays', () => {
      const fieldName1 = 'patchField1'
      const fieldName2 = 'patchArrayField1'

      const patch = {
        [fieldName1]: ['test', 'test2', { [fieldName2]: 'test3' }],
      }

      const paths = Marshaler.pathsFromPatch(patch)

      expect(paths).toBeDefined()
      expect(paths).toHaveLength(1)
      expect(paths).toEqual(expect.arrayContaining([fieldName1]))
    })

    it('should produce paths from deeply nested patch', () => {
      const fieldName1 = 'patchField1'
      const fieldName2 = 'nestedPatchField2'
      const fieldName3 = 'nestedPatchField3'
      const fieldName4 = 'nestedPatchField4'
      const fieldName5 = 'nestedPatchField5'
      const expected = [
        [fieldName1, fieldName2, fieldName3, fieldName4],
        [fieldName1, fieldName2, fieldName3, fieldName5],
      ]

      const patch = {
        [fieldName1]: {
          [fieldName2]: {
            [fieldName3]: {
              [fieldName4]: 'test',
              [fieldName5]: 'test',
            },
          },
        },
      }

      const paths = Marshaler.pathsFromPatch(patch)

      expect(paths).toBeDefined()
      expect(paths).toHaveLength(2)
      expect(paths).toEqual(expect.arrayContaining(expected))
    })
  })
})
