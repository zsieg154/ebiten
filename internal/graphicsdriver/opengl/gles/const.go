// Copyright 2020 The Ebiten Authors
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

// +build android ios

// Package gles implements Go bindings to OpenGL ES.
package gles

const (
	VERTEX_SHADER        = 0x8B31
	FRAGMENT_SHADER      = 0x8B30
	ARRAY_BUFFER         = 0x8892
	ELEMENT_ARRAY_BUFFER = 0x8893
	DYNAMIC_DRAW         = 0x88E8
	STREAM_DRAW          = 0x88E0
	PIXEL_UNPACK_BUFFER  = 0x88EC
	SHORT                = 0x1402
	FLOAT                = 0x1406

	ZERO                = 0
	ONE                 = 1
	SRC_ALPHA           = 0x0302
	DST_ALPHA           = 0x0304
	ONE_MINUS_SRC_ALPHA = 0x0303
	ONE_MINUS_DST_ALPHA = 0x0305
	DST_COLOR           = 0x0306

	FALSE = 0
	TRUE  = 1

	BLEND                = 0x0BE2
	CLAMP_TO_EDGE        = 0x812F
	COLOR_ATTACHMENT0    = 0x8CE0
	COMPILE_STATUS       = 0x8B81
	FRAMEBUFFER          = 0x8D40
	FRAMEBUFFER_BINDING  = 0x8CA6
	FRAMEBUFFER_COMPLETE = 0x8CD5
	HIGH_FLOAT           = 0x8DF2
	INFO_LOG_LENGTH      = 0x8B84
	LINK_STATUS          = 0x8B82
	MAX_TEXTURE_SIZE     = 0x0D33
	NEAREST              = 0x2600
	NO_ERROR             = 0
	READ_WRITE           = 0x88BA
	RGBA                 = 0x1908
	TEXTURE0             = 0x84C0
	TEXTURE_2D           = 0x0DE1
	TEXTURE_MAG_FILTER   = 0x2800
	TEXTURE_MIN_FILTER   = 0x2801
	TEXTURE_WRAP_S       = 0x2802
	TEXTURE_WRAP_T       = 0x2803
	TRIANGLES            = 0x0004
	UNPACK_ALIGNMENT     = 0x0CF5
	UNSIGNED_BYTE        = 0x1401
	UNSIGNED_SHORT       = 0x1403
	WRITE_ONLY           = 0x88B9
)
