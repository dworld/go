// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import (
	"unsafe"
)

func mapaccess1_fast32(t *maptype, h *hmap, key uint32) unsafe.Pointer {
	if raceenabled && h != nil {
		callerpc := gogetcallerpc(unsafe.Pointer(&t))
		fn := mapaccess1_fast32
		pc := **(**uintptr)(unsafe.Pointer(&fn))
		racereadpc(unsafe.Pointer(h), callerpc, pc)
	}
	if h == nil || h.count == 0 {
		return unsafe.Pointer(t.elem.zero)
	}
	var b *bmap
	if h.B == 0 {
		// One-bucket table.  No need to hash.
		b = (*bmap)(h.buckets)
	} else {
		hash := gohash(t.key.alg, unsafe.Pointer(&key), 4, uintptr(h.hash0))
		m := uintptr(1)<<h.B - 1
		b = (*bmap)(add(h.buckets, (hash&m)*uintptr(h.bucketsize)))
		if c := h.oldbuckets; c != nil {
			oldb := (*bmap)(add(c, (hash&(m>>1))*uintptr(h.bucketsize)))
			if !evacuated(oldb) {
				b = oldb
			}
		}
	}
	for {
		for i := uintptr(0); i < bucketCnt; i++ {
			k := *((*uint32)(add(unsafe.Pointer(b), dataOffset+i*4)))
			if k != key {
				continue
			}
			t := *((*uint8)(add(unsafe.Pointer(b), i))) // b.topbits[i] without the bounds check
			if t == empty {
				continue
			}
			return add(unsafe.Pointer(b), dataOffset+bucketCnt*4+i*uintptr(h.valuesize))
		}
		b = b.overflow
		if b == nil {
			return unsafe.Pointer(t.elem.zero)
		}
	}
}

func mapaccess2_fast32(t *maptype, h *hmap, key uint32) (unsafe.Pointer, bool) {
	if raceenabled && h != nil {
		callerpc := gogetcallerpc(unsafe.Pointer(&t))
		fn := mapaccess2_fast32
		pc := **(**uintptr)(unsafe.Pointer(&fn))
		racereadpc(unsafe.Pointer(h), callerpc, pc)
	}
	if h == nil || h.count == 0 {
		return unsafe.Pointer(t.elem.zero), false
	}
	var b *bmap
	if h.B == 0 {
		// One-bucket table.  No need to hash.
		b = (*bmap)(h.buckets)
	} else {
		hash := gohash(t.key.alg, unsafe.Pointer(&key), 4, uintptr(h.hash0))
		m := uintptr(1)<<h.B - 1
		b = (*bmap)(add(h.buckets, (hash&m)*uintptr(h.bucketsize)))
		if c := h.oldbuckets; c != nil {
			oldb := (*bmap)(add(c, (hash&(m>>1))*uintptr(h.bucketsize)))
			if !evacuated(oldb) {
				b = oldb
			}
		}
	}
	for {
		for i := uintptr(0); i < bucketCnt; i++ {
			k := *((*uint32)(add(unsafe.Pointer(b), dataOffset+i*4)))
			if k != key {
				continue
			}
			t := *((*uint8)(add(unsafe.Pointer(b), i))) // b.topbits[i] without the bounds check
			if t == empty {
				continue
			}
			return add(unsafe.Pointer(b), dataOffset+bucketCnt*4+i*uintptr(h.valuesize)), true
		}
		b = b.overflow
		if b == nil {
			return unsafe.Pointer(t.elem.zero), false
		}
	}
}

func mapaccess1_fast64(t *maptype, h *hmap, key uint64) unsafe.Pointer {
	if raceenabled && h != nil {
		callerpc := gogetcallerpc(unsafe.Pointer(&t))
		fn := mapaccess1_fast64
		pc := **(**uintptr)(unsafe.Pointer(&fn))
		racereadpc(unsafe.Pointer(h), callerpc, pc)
	}
	if h == nil || h.count == 0 {
		return unsafe.Pointer(t.elem.zero)
	}
	var b *bmap
	if h.B == 0 {
		// One-bucket table.  No need to hash.
		b = (*bmap)(h.buckets)
	} else {
		hash := gohash(t.key.alg, unsafe.Pointer(&key), 8, uintptr(h.hash0))
		m := uintptr(1)<<h.B - 1
		b = (*bmap)(add(h.buckets, (hash&m)*uintptr(h.bucketsize)))
		if c := h.oldbuckets; c != nil {
			oldb := (*bmap)(add(c, (hash&(m>>1))*uintptr(h.bucketsize)))
			if !evacuated(oldb) {
				b = oldb
			}
		}
	}
	for {
		for i := uintptr(0); i < bucketCnt; i++ {
			k := *((*uint64)(add(unsafe.Pointer(b), dataOffset+i*8)))
			if k != key {
				continue
			}
			t := *((*uint8)(add(unsafe.Pointer(b), i))) // b.topbits[i] without the bounds check
			if t == empty {
				continue
			}
			return add(unsafe.Pointer(b), dataOffset+bucketCnt*8+i*uintptr(h.valuesize))
		}
		b = b.overflow
		if b == nil {
			return unsafe.Pointer(t.elem.zero)
		}
	}
}

func mapaccess2_fast64(t *maptype, h *hmap, key uint64) (unsafe.Pointer, bool) {
	if raceenabled && h != nil {
		callerpc := gogetcallerpc(unsafe.Pointer(&t))
		fn := mapaccess2_fast64
		pc := **(**uintptr)(unsafe.Pointer(&fn))
		racereadpc(unsafe.Pointer(h), callerpc, pc)
	}
	if h == nil || h.count == 0 {
		return unsafe.Pointer(t.elem.zero), false
	}
	var b *bmap
	if h.B == 0 {
		// One-bucket table.  No need to hash.
		b = (*bmap)(h.buckets)
	} else {
		hash := gohash(t.key.alg, unsafe.Pointer(&key), 8, uintptr(h.hash0))
		m := uintptr(1)<<h.B - 1
		b = (*bmap)(add(h.buckets, (hash&m)*uintptr(h.bucketsize)))
		if c := h.oldbuckets; c != nil {
			oldb := (*bmap)(add(c, (hash&(m>>1))*uintptr(h.bucketsize)))
			if !evacuated(oldb) {
				b = oldb
			}
		}
	}
	for {
		for i := uintptr(0); i < bucketCnt; i++ {
			k := *((*uint64)(add(unsafe.Pointer(b), dataOffset+i*8)))
			if k != key {
				continue
			}
			t := *((*uint8)(add(unsafe.Pointer(b), i))) // b.topbits[i] without the bounds check
			if t == empty {
				continue
			}
			return add(unsafe.Pointer(b), dataOffset+bucketCnt*8+i*uintptr(h.valuesize)), true
		}
		b = b.overflow
		if b == nil {
			return unsafe.Pointer(t.elem.zero), false
		}
	}
}

func mapaccess1_faststr(t *maptype, h *hmap, ky string) unsafe.Pointer {
	if raceenabled && h != nil {
		callerpc := gogetcallerpc(unsafe.Pointer(&t))
		fn := mapaccess1_faststr
		pc := **(**uintptr)(unsafe.Pointer(&fn))
		racereadpc(unsafe.Pointer(h), callerpc, pc)
	}
	if h == nil || h.count == 0 {
		return unsafe.Pointer(t.elem.zero)
	}
	key := (*stringStruct)(unsafe.Pointer(&ky))
	if h.B == 0 {
		// One-bucket table.
		b := (*bmap)(h.buckets)
		if key.len < 32 {
			// short key, doing lots of comparisons is ok
			for i := uintptr(0); i < bucketCnt; i++ {
				t := *((*uint8)(add(unsafe.Pointer(b), i))) // b.topbits[i] without the bounds check
				if t == empty {
					continue
				}
				k := (*stringStruct)(add(unsafe.Pointer(b), dataOffset+i*2*ptrSize))
				if k.len != key.len {
					continue
				}
				if k.str == key.str || gomemeq(k.str, key.str, uintptr(key.len)) {
					return add(unsafe.Pointer(b), dataOffset+bucketCnt*2*ptrSize+i*uintptr(h.valuesize))
				}
			}
			return unsafe.Pointer(t.elem.zero)
		}
		// long key, try not to do more comparisons than necessary
		keymaybe := uintptr(bucketCnt)
		for i := uintptr(0); i < bucketCnt; i++ {
			t := *((*uint8)(add(unsafe.Pointer(b), i))) // b.topbits[i] without the bounds check
			if t == empty {
				continue
			}
			k := (*stringStruct)(add(unsafe.Pointer(b), dataOffset+i*2*ptrSize))
			if k.len != key.len {
				continue
			}
			if k.str == key.str {
				return add(unsafe.Pointer(b), dataOffset+bucketCnt*2*ptrSize+i*uintptr(h.valuesize))
			}
			// check first 4 bytes
			// TODO: on amd64/386 at least, make this compile to one 4-byte comparison instead of
			// four 1-byte comparisons.
			if *((*[4]byte)(key.str)) != *((*[4]byte)(k.str)) {
				continue
			}
			// check last 4 bytes
			if *((*[4]byte)(add(key.str, uintptr(key.len)-4))) != *((*[4]byte)(add(k.str, uintptr(key.len)-4))) {
				continue
			}
			if keymaybe != bucketCnt {
				// Two keys are potential matches.  Use hash to distinguish them.
				goto dohash
			}
			keymaybe = i
		}
		if keymaybe != bucketCnt {
			k := (*stringStruct)(add(unsafe.Pointer(b), dataOffset+keymaybe*2*ptrSize))
			if gomemeq(k.str, key.str, uintptr(key.len)) {
				return add(unsafe.Pointer(b), dataOffset+bucketCnt*2*ptrSize+keymaybe*uintptr(h.valuesize))
			}
		}
		return unsafe.Pointer(t.elem.zero)
	}
dohash:
	hash := gohash(t.key.alg, unsafe.Pointer(&ky), 2*ptrSize, uintptr(h.hash0))
	m := uintptr(1)<<h.B - 1
	b := (*bmap)(add(h.buckets, (hash&m)*uintptr(h.bucketsize)))
	if c := h.oldbuckets; c != nil {
		oldb := (*bmap)(add(c, (hash&(m>>1))*uintptr(h.bucketsize)))
		if !evacuated(oldb) {
			b = oldb
		}
	}
	top := uint8(hash >> (ptrSize*8 - 8))
	if top < minTopHash {
		top += minTopHash
	}
	for {
		for i := uintptr(0); i < bucketCnt; i++ {
			t := *((*uint8)(add(unsafe.Pointer(b), i))) // b.topbits[i] without the bounds check
			if t != top {
				continue
			}
			k := (*stringStruct)(add(unsafe.Pointer(b), dataOffset+i*2*ptrSize))
			if k.len != key.len {
				continue
			}
			if k.str == key.str || gomemeq(k.str, key.str, uintptr(key.len)) {
				return add(unsafe.Pointer(b), dataOffset+bucketCnt*2*ptrSize+i*uintptr(h.valuesize))
			}
		}
		b = b.overflow
		if b == nil {
			return unsafe.Pointer(t.elem.zero)
		}
	}
}

func mapaccess2_faststr(t *maptype, h *hmap, ky string) (unsafe.Pointer, bool) {
	if raceenabled && h != nil {
		callerpc := gogetcallerpc(unsafe.Pointer(&t))
		fn := mapaccess2_faststr
		pc := **(**uintptr)(unsafe.Pointer(&fn))
		racereadpc(unsafe.Pointer(h), callerpc, pc)
	}
	if h == nil || h.count == 0 {
		return unsafe.Pointer(t.elem.zero), false
	}
	key := (*stringStruct)(unsafe.Pointer(&ky))
	if h.B == 0 {
		// One-bucket table.
		b := (*bmap)(h.buckets)
		if key.len < 32 {
			// short key, doing lots of comparisons is ok
			for i := uintptr(0); i < bucketCnt; i++ {
				t := *((*uint8)(add(unsafe.Pointer(b), i))) // b.topbits[i] without the bounds check
				if t == empty {
					continue
				}
				k := (*stringStruct)(add(unsafe.Pointer(b), dataOffset+i*2*ptrSize))
				if k.len != key.len {
					continue
				}
				if k.str == key.str || gomemeq(k.str, key.str, uintptr(key.len)) {
					return add(unsafe.Pointer(b), dataOffset+bucketCnt*2*ptrSize+i*uintptr(h.valuesize)), true
				}
			}
			return unsafe.Pointer(t.elem.zero), false
		}
		// long key, try not to do more comparisons than necessary
		keymaybe := uintptr(bucketCnt)
		for i := uintptr(0); i < bucketCnt; i++ {
			t := *((*uint8)(add(unsafe.Pointer(b), i))) // b.topbits[i] without the bounds check
			if t == empty {
				continue
			}
			k := (*stringStruct)(add(unsafe.Pointer(b), dataOffset+i*2*ptrSize))
			if k.len != key.len {
				continue
			}
			if k.str == key.str {
				return add(unsafe.Pointer(b), dataOffset+bucketCnt*2*ptrSize+i*uintptr(h.valuesize)), true
			}
			// check first 4 bytes
			if *((*[4]byte)(key.str)) != *((*[4]byte)(k.str)) {
				continue
			}
			// check last 4 bytes
			if *((*[4]byte)(add(key.str, uintptr(key.len)-4))) != *((*[4]byte)(add(k.str, uintptr(key.len)-4))) {
				continue
			}
			if keymaybe != bucketCnt {
				// Two keys are potential matches.  Use hash to distinguish them.
				goto dohash
			}
			keymaybe = i
		}
		if keymaybe != bucketCnt {
			k := (*stringStruct)(add(unsafe.Pointer(b), dataOffset+keymaybe*2*ptrSize))
			if gomemeq(k.str, key.str, uintptr(key.len)) {
				return add(unsafe.Pointer(b), dataOffset+bucketCnt*2*ptrSize+keymaybe*uintptr(h.valuesize)), true
			}
		}
		return unsafe.Pointer(t.elem.zero), false
	}
dohash:
	hash := gohash(t.key.alg, unsafe.Pointer(&ky), 2*ptrSize, uintptr(h.hash0))
	m := uintptr(1)<<h.B - 1
	b := (*bmap)(add(h.buckets, (hash&m)*uintptr(h.bucketsize)))
	if c := h.oldbuckets; c != nil {
		oldb := (*bmap)(add(c, (hash&(m>>1))*uintptr(h.bucketsize)))
		if !evacuated(oldb) {
			b = oldb
		}
	}
	top := uint8(hash >> (ptrSize*8 - 8))
	if top < minTopHash {
		top += minTopHash
	}
	for {
		for i := uintptr(0); i < bucketCnt; i++ {
			t := *((*uint8)(add(unsafe.Pointer(b), i))) // b.topbits[i] without the bounds check
			if t != top {
				continue
			}
			k := (*stringStruct)(add(unsafe.Pointer(b), dataOffset+i*2*ptrSize))
			if k.len != key.len {
				continue
			}
			if k.str == key.str || gomemeq(k.str, key.str, uintptr(key.len)) {
				return add(unsafe.Pointer(b), dataOffset+bucketCnt*2*ptrSize+i*uintptr(h.valuesize)), true
			}
		}
		b = b.overflow
		if b == nil {
			return unsafe.Pointer(t.elem.zero), false
		}
	}
}