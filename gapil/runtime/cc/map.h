// Copyright (C) 2018 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

#ifndef __GAPIL_RUNTIME_MAP_H__
#define __GAPIL_RUNTIME_MAP_H__

#include <stddef.h>
#include <stdint.h>

#include "maker.h"

#define GAPIL_MAP_ELEMENT_EMPTY 0
#define GAPIL_MAP_ELEMENT_FULL 1
#define GAPIL_MAP_ELEMENT_USED 2

#define GAPIL_MAP_GROW_MULTIPLIER 4
#define GAPIL_MIN_MAP_SIZE 32
#define GAPIL_MAP_MAX_CAPACITY 0.8f

namespace core {
class Arena;
}  // namespace core

namespace gapil {

// Map is an associative container that is compatible with the maps produced by
// the gapil compiler. Maps hold references to their data, and several maps may
// share the same underlying data.
// Search, insertion, and removal of elements are O(1).
// Like std::unordered_map, elements are stored in no particular order.
template <typename K, typename V, bool DENSE>
class Map {
 private:
  struct Allocation;

 public:
  struct element {
    uint64_t used;
    K first;
    V second;
  };

  using key_type = K;
  using value_type = V;

  class iterator {
   public:
    inline iterator(const iterator& it);

    inline bool operator==(const iterator& other);

    inline bool operator!=(const iterator& other);

    inline element& operator*();

    inline element* operator->();

    inline const iterator& operator++();

    inline iterator operator++(int);

   private:
    friend class Map<K, V, DENSE>;
    element* elem;
    Allocation* map;

    inline iterator(element* elem, Allocation* map);
  };

  class const_iterator {
   public:
    inline const_iterator(const iterator& it);

    inline bool operator==(const const_iterator& other);

    inline bool operator!=(const const_iterator& other);

    inline const element& operator*();

    inline const element* operator->();

    inline const_iterator& operator++();

    inline const_iterator operator++(int);

   private:
    friend class Map<K, V, DENSE>;
    const element* elem;
    const Allocation* map;

    inline const_iterator(const element* elem, const Allocation* map);
  };

  // Constructs a new empty map.
  Map(core::Arena*);

  // Constructs a map which shares the data with other.
  Map(const Map<K, V, DENSE>& other);

  Map(Map<K, V, DENSE>&&);
  ~Map();

  // Returns a copy of this map with the same elements.
  Map<K, V, DENSE> clone() const;

  // Makes this map refer to the RHS map.
  Map<K, V, DENSE>& operator=(const Map<K, V, DENSE>&);

  // Returns the arena this map belongs to.
  inline core::Arena* arena() const;

  // Returns the number of elements that can be held in the map before
  // Re-allocating the internal buffer.
  inline uint64_t capacity() const;

  // Returns the number of elements currently held in the map.
  inline uint64_t count() const;

  // Returns true if the map has no elements.
  inline bool empty() const;

  // Returns true if the map contains the given key.
  inline bool contains(const K&) const __attribute__((always_inline));

  // Returns a constant iterator to the beginning.
  inline const const_iterator begin() const;

  // Returns an iterator to the beginning.
  inline iterator begin();

  // Returns a constant iterator to the end.
  inline const_iterator end() const;

  // Returns an iterator to the end.
  inline iterator end();

  // Removes the element with the given key from the map.
  inline void erase(const K& k);

  // Removes the element with the given key from the map.
  inline void erase(const_iterator it);

  // Clears the map, but does not free any of the memory.
  inline void clear();

  // Returns a reference to the element with the given key, creating and
  // adding an element if one did not exist.
  template <typename T>
  inline V& operator[](const T& key);

  // Returns an iterator to an element with the given key.
  // If no such element is found, the end iterator is returned.
  inline iterator find(const K& key);

  // Returns an iterator to an element with the given key.
  // If no such element is found, the end iterator is returned.
  inline const_iterator find(const K& k) const;

  // finds a key in the map and returns the value. If no value is present
  // Returns the zero for that type.
  inline V findOrZero(const K& key) const;

  // instance_ptr returns an opaque pointer to the underlying map data.
  // This can be used for map equality.
  inline const void* instance_ptr() const;

 private:
  // The shared data of this map.
  struct Allocation {
    uint32_t ref_count;  // number of owners of this map.
    core::Arena* arena;  // arena that owns this allocation and elements buffer.
    uint64_t count;      // number of elements in the map.
    uint64_t capacity;   // size of the elements buffer.
    void* elements;      // pointer to the elements buffer.

    // The following methods must match those generated by the gapil compiler.
    bool contains(K);

    inline V* index(K, bool) __attribute__((always_inline));
    V lookup(K);
    void remove(K);
    void clear_for_delete();
    void clear_keep();
    void reference();
    void release();

    inline const element* els() const;
    inline element* els();
  };

  struct SVOAllocation {
    Allocation alloc;
    element els[GAPIL_MIN_MAP_SIZE];
  };

  Allocation* ptr;
};

////////////////////////////////////////////////////////////////////////////////
// Map<K, V, DENSE>::iterator //
////////////////////////////////////////////////////////////////////////////////

template <typename K, typename V, bool DENSE>
Map<K, V, DENSE>::iterator::iterator(element* elem, Allocation* map)
    : elem(elem), map(map) {}

template <typename K, typename V, bool DENSE>
Map<K, V, DENSE>::iterator::iterator(const iterator& it)
    : elem(it.elem), map(it.map) {}

template <typename K, typename V, bool DENSE>
bool Map<K, V, DENSE>::iterator::operator==(const iterator& other) {
  return map == other.map && elem == other.elem;
}

template <typename K, typename V, bool DENSE>
bool Map<K, V, DENSE>::iterator::operator!=(const iterator& other) {
  return !(*this == other);
}

template <typename K, typename V, bool DENSE>
typename Map<K, V, DENSE>::element& Map<K, V, DENSE>::iterator::operator*() {
  return *elem;
}

template <typename K, typename V, bool DENSE>
typename Map<K, V, DENSE>::element* Map<K, V, DENSE>::iterator::operator->() {
  return elem;
}

template <typename K, typename V, bool DENSE>
const typename Map<K, V, DENSE>::iterator&
Map<K, V, DENSE>::iterator::operator++() {
  size_t offset = elem - reinterpret_cast<element*>(map->elements);
  for (size_t i = offset; i < map->capacity; ++i) {
    ++elem;
    if (i == map->capacity - 1) {
      break;
    }
    if (elem->used == GAPIL_MAP_ELEMENT_FULL) {
      break;
    }
  }
  return *this;
}

template <typename K, typename V, bool DENSE>
typename Map<K, V, DENSE>::iterator Map<K, V, DENSE>::iterator::operator++(
    int) {
  iterator ret = *this;
  ++(*this);
  return ret;
}

////////////////////////////////////////////////////////////////////////////////
// Map<K, V, DENSE>::const_iterator //
////////////////////////////////////////////////////////////////////////////////

template <typename K, typename V, bool DENSE>
Map<K, V, DENSE>::const_iterator::const_iterator(const element* elem,
                                                 const Allocation* map)
    : elem(elem), map(map) {}

template <typename K, typename V, bool DENSE>
Map<K, V, DENSE>::const_iterator::const_iterator(const iterator& it)
    : elem(it.elem), map(it.map) {}

template <typename K, typename V, bool DENSE>
bool Map<K, V, DENSE>::const_iterator::operator==(const const_iterator& other) {
  return map == other.map && elem == other.elem;
}

template <typename K, typename V, bool DENSE>
bool Map<K, V, DENSE>::const_iterator::operator!=(const const_iterator& other) {
  return !(*this == other);
}

template <typename K, typename V, bool DENSE>
const typename Map<K, V, DENSE>::element&
Map<K, V, DENSE>::const_iterator::operator*() {
  return *elem;
}

template <typename K, typename V, bool DENSE>
const typename Map<K, V, DENSE>::element*
Map<K, V, DENSE>::const_iterator::operator->() {
  return elem;
}

template <typename K, typename V, bool DENSE>
typename Map<K, V, DENSE>::const_iterator&
Map<K, V, DENSE>::const_iterator::operator++() {
  size_t offset = elem - reinterpret_cast<element*>(map->elements);
  for (size_t i = offset; i < map->capacity; ++i) {
    ++elem;
    if (i == map->capacity - 1) {
      break;
    }
    if (elem->used == GAPIL_MAP_ELEMENT_FULL) {
      break;
    }
  }
  return *this;
}

template <typename K, typename V, bool DENSE>
typename Map<K, V, DENSE>::const_iterator
Map<K, V, DENSE>::const_iterator::operator++(int) {
  const_iterator ret = *this;
  ++(*this);
  return ret;
}

////////////////////////////////////////////////////////////////////////////////
// Map<K, V, DENSE> //
////////////////////////////////////////////////////////////////////////////////

template <typename K, typename V, bool DENSE>
core::Arena* Map<K, V, DENSE>::arena() const {
  return reinterpret_cast<core::Arena*>(ptr->arena);
}

template <typename K, typename V, bool DENSE>
uint64_t Map<K, V, DENSE>::capacity() const {
  return ptr->capacity;
}

template <typename K, typename V, bool DENSE>
uint64_t Map<K, V, DENSE>::count() const {
  return ptr->count;
}

template <typename K, typename V, bool DENSE>
bool Map<K, V, DENSE>::empty() const {
  return ptr->count == 0;
}

template <typename K, typename V, bool DENSE>
bool Map<K, V, DENSE>::contains(const K& key) const {
  return ptr->contains(key);
}

template <typename K, typename V, bool DENSE>
const typename Map<K, V, DENSE>::const_iterator Map<K, V, DENSE>::begin()
    const {
  auto it = const_iterator{ptr->els(), ptr};
  for (size_t i = 0; i < ptr->capacity; ++i) {
    if (it.elem->used == GAPIL_MAP_ELEMENT_FULL) {
      break;
    }
    it.elem++;
  }
  return it;
}

template <typename K, typename V, bool DENSE>
typename Map<K, V, DENSE>::iterator Map<K, V, DENSE>::begin() {
  auto it = iterator{ptr->els(), ptr};
  for (size_t i = 0; i < ptr->capacity; ++i) {
    if (it.elem->used == GAPIL_MAP_ELEMENT_FULL) {
      break;
    }
    it.elem++;
  }
  return it;
}

template <typename K, typename V, bool DENSE>
typename Map<K, V, DENSE>::iterator Map<K, V, DENSE>::end() {
  return iterator{ptr->els() + capacity(), ptr};
}

template <typename K, typename V, bool DENSE>
typename Map<K, V, DENSE>::const_iterator Map<K, V, DENSE>::end() const {
  return const_iterator{ptr->els() + capacity(), ptr};
}

template <typename K, typename V, bool DENSE>
void Map<K, V, DENSE>::erase(const K& k) {
  ptr->remove(k);
}

template <typename K, typename V, bool DENSE>
void Map<K, V, DENSE>::erase(const_iterator it) {
  ptr->remove(it->first);
}

template <typename K, typename V, bool DENSE>
void Map<K, V, DENSE>::clear() {
  ptr->clear_keep();
}

template <typename K, typename V, bool DENSE>
template <typename T>
V& Map<K, V, DENSE>::operator[](const T& key) {
  V* v = ptr->index(key, true);
  return *v;
}

template <typename K, typename V, bool DENSE>
typename Map<K, V, DENSE>::iterator Map<K, V, DENSE>::find(const K& key) {
  V* idx = ptr->index(key, false);
  if (idx == nullptr) {
    return end();
  }
  size_t offs = (reinterpret_cast<uintptr_t>(idx) -
                 reinterpret_cast<uintptr_t>(ptr->els())) /
                sizeof(element);
  return iterator{ptr->els() + offs, ptr};
}

template <typename K, typename V, bool DENSE>
typename Map<K, V, DENSE>::const_iterator Map<K, V, DENSE>::find(
    const K& k) const {
  const V* idx = ptr->index(k, false);
  if (idx == nullptr) {
    return end();
  }
  size_t offs = (reinterpret_cast<uintptr_t>(idx) -
                 reinterpret_cast<uintptr_t>(ptr->els())) /
                sizeof(element);
  return const_iterator{ptr->els() + offs, ptr};
}

template <typename K, typename V, bool DENSE>
V Map<K, V, DENSE>::findOrZero(const K& key) const {
  auto it = find(key);
  if (it == end()) {
    auto arena = reinterpret_cast<core::Arena*>(ptr->arena);
    return make<V>(arena);
  }
  return it->second;
}

template <typename K, typename V, bool DENSE>
inline const void* Map<K, V, DENSE>::instance_ptr() const {
  return ptr;
}

////////////////////////////////////////////////////////////////////////////////
// Map<K, V, DENSE>::Allocation //
////////////////////////////////////////////////////////////////////////////////

template <typename K, typename V, bool DENSE>
const typename Map<K, V, DENSE>::element* Map<K, V, DENSE>::Allocation::els()
    const {
  return reinterpret_cast<const element*>(elements);
}

template <typename K, typename V, bool DENSE>
typename Map<K, V, DENSE>::element* Map<K, V, DENSE>::Allocation::els() {
  return reinterpret_cast<element*>(elements);
}

}  // namespace gapil

#endif  // __GAPIL_RUNTIME_MAP_H__
