package zscratchpad

// ~/go/bin/gencode go -schema ./sources/lib/gencode.schema -out ./sources/lib/gencode.go -package zscratchpad -unsafe

import (
	"io"
	"time"
	"unsafe"
)

var (
	_ = unsafe.Sizeof(0)
	_ = io.ReadFull
	_ = time.Now()
)

/*
type Document struct {
	Identifier        string
	Library           string
	Path              string
	PathInLibrary     string
	Title             string
	TitleAlternatives []string
	SourceFingerprint string
	Format            string
	BodyLines         []string
	BodyEmpty         bool
	BodyFingerprint   string
	EditEnabled       bool
	Timestamp         time.Time
}
*/

func (d *Document) Size() (s uint64) {

	{
		l := uint64(len(d.Identifier))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	{
		l := uint64(len(d.Library))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	{
		l := uint64(len(d.Path))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	{
		l := uint64(len(d.PathInLibrary))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	{
		l := uint64(len(d.Title))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	{
		l := uint64(len(d.TitleAlternatives))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}

		for k0 := range d.TitleAlternatives {

			{
				l := uint64(len(d.TitleAlternatives[k0]))

				{

					t := l
					for t >= 0x80 {
						t >>= 7
						s++
					}
					s++

				}
				s += l
			}

		}

	}
	{
		l := uint64(len(d.SourceFingerprint))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	{
		l := uint64(len(d.Format))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	{
		l := uint64(len(d.BodyLines))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}

		for k0 := range d.BodyLines {

			{
				l := uint64(len(d.BodyLines[k0]))

				{

					t := l
					for t >= 0x80 {
						t >>= 7
						s++
					}
					s++

				}
				s += l
			}

		}

	}
	{
		l := uint64(len(d.BodyFingerprint))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	s += 17
	return
}
func (d *Document) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{
		l := uint64(len(d.Identifier))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+0] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+0] = byte(t)
			i++

		}
		copy(buf[i+0:], d.Identifier)
		i += l
	}
	{
		l := uint64(len(d.Library))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+0] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+0] = byte(t)
			i++

		}
		copy(buf[i+0:], d.Library)
		i += l
	}
	{
		l := uint64(len(d.Path))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+0] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+0] = byte(t)
			i++

		}
		copy(buf[i+0:], d.Path)
		i += l
	}
	{
		l := uint64(len(d.PathInLibrary))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+0] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+0] = byte(t)
			i++

		}
		copy(buf[i+0:], d.PathInLibrary)
		i += l
	}
	{
		l := uint64(len(d.Title))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+0] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+0] = byte(t)
			i++

		}
		copy(buf[i+0:], d.Title)
		i += l
	}
	{
		l := uint64(len(d.TitleAlternatives))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+0] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+0] = byte(t)
			i++

		}
		for k0 := range d.TitleAlternatives {

			{
				l := uint64(len(d.TitleAlternatives[k0]))

				{

					t := uint64(l)

					for t >= 0x80 {
						buf[i+0] = byte(t) | 0x80
						t >>= 7
						i++
					}
					buf[i+0] = byte(t)
					i++

				}
				copy(buf[i+0:], d.TitleAlternatives[k0])
				i += l
			}

		}
	}
	{
		l := uint64(len(d.SourceFingerprint))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+0] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+0] = byte(t)
			i++

		}
		copy(buf[i+0:], d.SourceFingerprint)
		i += l
	}
	{
		l := uint64(len(d.Format))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+0] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+0] = byte(t)
			i++

		}
		copy(buf[i+0:], d.Format)
		i += l
	}
	{
		l := uint64(len(d.BodyLines))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+0] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+0] = byte(t)
			i++

		}
		for k0 := range d.BodyLines {

			{
				l := uint64(len(d.BodyLines[k0]))

				{

					t := uint64(l)

					for t >= 0x80 {
						buf[i+0] = byte(t) | 0x80
						t >>= 7
						i++
					}
					buf[i+0] = byte(t)
					i++

				}
				copy(buf[i+0:], d.BodyLines[k0])
				i += l
			}

		}
	}
	{
		if d.BodyEmpty {
			buf[i+0] = 1
		} else {
			buf[i+0] = 0
		}
	}
	{
		l := uint64(len(d.BodyFingerprint))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+1] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+1] = byte(t)
			i++

		}
		copy(buf[i+1:], d.BodyFingerprint)
		i += l
	}
	{
		if d.EditEnabled {
			buf[i+1] = 1
		} else {
			buf[i+1] = 0
		}
	}
	{
		b, err := d.Timestamp.MarshalBinary()
		if err != nil {
			return nil, err
		}
		copy(buf[i+2:], b)
	}
	return buf[:i+17], nil
}

func (d *Document) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i+0]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.Identifier = string(buf[i+0 : i+0+l])
		i += l
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i+0]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.Library = string(buf[i+0 : i+0+l])
		i += l
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i+0]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.Path = string(buf[i+0 : i+0+l])
		i += l
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i+0]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.PathInLibrary = string(buf[i+0 : i+0+l])
		i += l
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i+0]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.Title = string(buf[i+0 : i+0+l])
		i += l
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i+0]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		if uint64(cap(d.TitleAlternatives)) >= l {
			d.TitleAlternatives = d.TitleAlternatives[:l]
		} else {
			d.TitleAlternatives = make([]string, l)
		}
		for k0 := range d.TitleAlternatives {

			{
				l := uint64(0)

				{

					bs := uint8(7)
					t := uint64(buf[i+0] & 0x7F)
					for buf[i+0]&0x80 == 0x80 {
						i++
						t |= uint64(buf[i+0]&0x7F) << bs
						bs += 7
					}
					i++

					l = t

				}
				d.TitleAlternatives[k0] = string(buf[i+0 : i+0+l])
				i += l
			}

		}
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i+0]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.SourceFingerprint = string(buf[i+0 : i+0+l])
		i += l
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i+0]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.Format = string(buf[i+0 : i+0+l])
		i += l
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i+0]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		if uint64(cap(d.BodyLines)) >= l {
			d.BodyLines = d.BodyLines[:l]
		} else {
			d.BodyLines = make([]string, l)
		}
		for k0 := range d.BodyLines {

			{
				l := uint64(0)

				{

					bs := uint8(7)
					t := uint64(buf[i+0] & 0x7F)
					for buf[i+0]&0x80 == 0x80 {
						i++
						t |= uint64(buf[i+0]&0x7F) << bs
						bs += 7
					}
					i++

					l = t

				}
				d.BodyLines[k0] = string(buf[i+0 : i+0+l])
				i += l
			}

		}
	}
	{
		d.BodyEmpty = buf[i+0] == 1
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+1] & 0x7F)
			for buf[i+1]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+1]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.BodyFingerprint = string(buf[i+1 : i+1+l])
		i += l
	}
	{
		d.EditEnabled = buf[i+1] == 1
	}
	{
		d.Timestamp.UnmarshalBinary(buf[i+2 : i+2+15])
	}
	return i + 17, nil
}

/*
type Library struct {
	Identifier                     string
	Name                           string
	Paths                          []string
	Disabled                       bool
	EditEnabled                    bool
	CreateEnabled                  bool
	CreatePath                     string
	CreateNameTimestampLength      uint8
	CreateNameRandomLength         uint8
	CreateExtension                string
	SnapshotEnabled                bool
	SnapshotExtension              string
	IncludeGlobPatterns            []string
	ExcludeGlobPatterns            []string
	IncludeRegexPatterns           []string
	ExcludeRegexPatterns           []string
	UseTitlePrefix                 string
	UseLibraryAsIdentifierPrefix   bool
	UsePathInLibraryAsIdentifier   bool
	UseFileNameAsIdentifier        bool
	UsePathFingerprintAsIdentifier bool
	UseFileExtensionAsFormat       bool
}
*/

func (d *Library) Size() (s uint64) {

	{
		l := uint64(len(d.Identifier))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	{
		l := uint64(len(d.Name))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	{
		l := uint64(len(d.Paths))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}

		for k0 := range d.Paths {

			{
				l := uint64(len(d.Paths[k0]))

				{

					t := l
					for t >= 0x80 {
						t >>= 7
						s++
					}
					s++

				}
				s += l
			}

		}

	}
	{
		l := uint64(len(d.CreatePath))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	{
		l := uint64(len(d.CreateExtension))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	{
		l := uint64(len(d.SnapshotExtension))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	{
		l := uint64(len(d.IncludeGlobPatterns))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}

		for k0 := range d.IncludeGlobPatterns {

			{
				l := uint64(len(d.IncludeGlobPatterns[k0]))

				{

					t := l
					for t >= 0x80 {
						t >>= 7
						s++
					}
					s++

				}
				s += l
			}

		}

	}
	{
		l := uint64(len(d.ExcludeGlobPatterns))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}

		for k0 := range d.ExcludeGlobPatterns {

			{
				l := uint64(len(d.ExcludeGlobPatterns[k0]))

				{

					t := l
					for t >= 0x80 {
						t >>= 7
						s++
					}
					s++

				}
				s += l
			}

		}

	}
	{
		l := uint64(len(d.IncludeRegexPatterns))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}

		for k0 := range d.IncludeRegexPatterns {

			{
				l := uint64(len(d.IncludeRegexPatterns[k0]))

				{

					t := l
					for t >= 0x80 {
						t >>= 7
						s++
					}
					s++

				}
				s += l
			}

		}

	}
	{
		l := uint64(len(d.ExcludeRegexPatterns))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}

		for k0 := range d.ExcludeRegexPatterns {

			{
				l := uint64(len(d.ExcludeRegexPatterns[k0]))

				{

					t := l
					for t >= 0x80 {
						t >>= 7
						s++
					}
					s++

				}
				s += l
			}

		}

	}
	{
		l := uint64(len(d.UseTitlePrefix))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	s += 11
	return
}
func (d *Library) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{
		l := uint64(len(d.Identifier))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+0] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+0] = byte(t)
			i++

		}
		copy(buf[i+0:], d.Identifier)
		i += l
	}
	{
		l := uint64(len(d.Name))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+0] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+0] = byte(t)
			i++

		}
		copy(buf[i+0:], d.Name)
		i += l
	}
	{
		l := uint64(len(d.Paths))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+0] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+0] = byte(t)
			i++

		}
		for k0 := range d.Paths {

			{
				l := uint64(len(d.Paths[k0]))

				{

					t := uint64(l)

					for t >= 0x80 {
						buf[i+0] = byte(t) | 0x80
						t >>= 7
						i++
					}
					buf[i+0] = byte(t)
					i++

				}
				copy(buf[i+0:], d.Paths[k0])
				i += l
			}

		}
	}
	{
		if d.Disabled {
			buf[i+0] = 1
		} else {
			buf[i+0] = 0
		}
	}
	{
		if d.EditEnabled {
			buf[i+1] = 1
		} else {
			buf[i+1] = 0
		}
	}
	{
		if d.CreateEnabled {
			buf[i+2] = 1
		} else {
			buf[i+2] = 0
		}
	}
	{
		l := uint64(len(d.CreatePath))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+3] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+3] = byte(t)
			i++

		}
		copy(buf[i+3:], d.CreatePath)
		i += l
	}
	{

		*(*uint8)(unsafe.Pointer(&buf[i+3])) = d.CreateNameTimestampLength

	}
	{

		*(*uint8)(unsafe.Pointer(&buf[i+4])) = d.CreateNameRandomLength

	}
	{
		l := uint64(len(d.CreateExtension))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+5] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+5] = byte(t)
			i++

		}
		copy(buf[i+5:], d.CreateExtension)
		i += l
	}
	{
		if d.SnapshotEnabled {
			buf[i+5] = 1
		} else {
			buf[i+5] = 0
		}
	}
	{
		l := uint64(len(d.SnapshotExtension))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+6] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+6] = byte(t)
			i++

		}
		copy(buf[i+6:], d.SnapshotExtension)
		i += l
	}
	{
		l := uint64(len(d.IncludeGlobPatterns))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+6] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+6] = byte(t)
			i++

		}
		for k0 := range d.IncludeGlobPatterns {

			{
				l := uint64(len(d.IncludeGlobPatterns[k0]))

				{

					t := uint64(l)

					for t >= 0x80 {
						buf[i+6] = byte(t) | 0x80
						t >>= 7
						i++
					}
					buf[i+6] = byte(t)
					i++

				}
				copy(buf[i+6:], d.IncludeGlobPatterns[k0])
				i += l
			}

		}
	}
	{
		l := uint64(len(d.ExcludeGlobPatterns))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+6] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+6] = byte(t)
			i++

		}
		for k0 := range d.ExcludeGlobPatterns {

			{
				l := uint64(len(d.ExcludeGlobPatterns[k0]))

				{

					t := uint64(l)

					for t >= 0x80 {
						buf[i+6] = byte(t) | 0x80
						t >>= 7
						i++
					}
					buf[i+6] = byte(t)
					i++

				}
				copy(buf[i+6:], d.ExcludeGlobPatterns[k0])
				i += l
			}

		}
	}
	{
		l := uint64(len(d.IncludeRegexPatterns))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+6] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+6] = byte(t)
			i++

		}
		for k0 := range d.IncludeRegexPatterns {

			{
				l := uint64(len(d.IncludeRegexPatterns[k0]))

				{

					t := uint64(l)

					for t >= 0x80 {
						buf[i+6] = byte(t) | 0x80
						t >>= 7
						i++
					}
					buf[i+6] = byte(t)
					i++

				}
				copy(buf[i+6:], d.IncludeRegexPatterns[k0])
				i += l
			}

		}
	}
	{
		l := uint64(len(d.ExcludeRegexPatterns))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+6] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+6] = byte(t)
			i++

		}
		for k0 := range d.ExcludeRegexPatterns {

			{
				l := uint64(len(d.ExcludeRegexPatterns[k0]))

				{

					t := uint64(l)

					for t >= 0x80 {
						buf[i+6] = byte(t) | 0x80
						t >>= 7
						i++
					}
					buf[i+6] = byte(t)
					i++

				}
				copy(buf[i+6:], d.ExcludeRegexPatterns[k0])
				i += l
			}

		}
	}
	{
		l := uint64(len(d.UseTitlePrefix))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+6] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+6] = byte(t)
			i++

		}
		copy(buf[i+6:], d.UseTitlePrefix)
		i += l
	}
	{
		if d.UseLibraryAsIdentifierPrefix {
			buf[i+6] = 1
		} else {
			buf[i+6] = 0
		}
	}
	{
		if d.UsePathInLibraryAsIdentifier {
			buf[i+7] = 1
		} else {
			buf[i+7] = 0
		}
	}
	{
		if d.UseFileNameAsIdentifier {
			buf[i+8] = 1
		} else {
			buf[i+8] = 0
		}
	}
	{
		if d.UsePathFingerprintAsIdentifier {
			buf[i+9] = 1
		} else {
			buf[i+9] = 0
		}
	}
	{
		if d.UseFileExtensionAsFormat {
			buf[i+10] = 1
		} else {
			buf[i+10] = 0
		}
	}
	return buf[:i+11], nil
}

func (d *Library) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i+0]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.Identifier = string(buf[i+0 : i+0+l])
		i += l
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i+0]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.Name = string(buf[i+0 : i+0+l])
		i += l
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i+0]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		if uint64(cap(d.Paths)) >= l {
			d.Paths = d.Paths[:l]
		} else {
			d.Paths = make([]string, l)
		}
		for k0 := range d.Paths {

			{
				l := uint64(0)

				{

					bs := uint8(7)
					t := uint64(buf[i+0] & 0x7F)
					for buf[i+0]&0x80 == 0x80 {
						i++
						t |= uint64(buf[i+0]&0x7F) << bs
						bs += 7
					}
					i++

					l = t

				}
				d.Paths[k0] = string(buf[i+0 : i+0+l])
				i += l
			}

		}
	}
	{
		d.Disabled = buf[i+0] == 1
	}
	{
		d.EditEnabled = buf[i+1] == 1
	}
	{
		d.CreateEnabled = buf[i+2] == 1
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+3] & 0x7F)
			for buf[i+3]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+3]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.CreatePath = string(buf[i+3 : i+3+l])
		i += l
	}
	{

		d.CreateNameTimestampLength = *(*uint8)(unsafe.Pointer(&buf[i+3]))

	}
	{

		d.CreateNameRandomLength = *(*uint8)(unsafe.Pointer(&buf[i+4]))

	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+5] & 0x7F)
			for buf[i+5]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+5]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.CreateExtension = string(buf[i+5 : i+5+l])
		i += l
	}
	{
		d.SnapshotEnabled = buf[i+5] == 1
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+6] & 0x7F)
			for buf[i+6]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+6]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.SnapshotExtension = string(buf[i+6 : i+6+l])
		i += l
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+6] & 0x7F)
			for buf[i+6]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+6]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		if uint64(cap(d.IncludeGlobPatterns)) >= l {
			d.IncludeGlobPatterns = d.IncludeGlobPatterns[:l]
		} else {
			d.IncludeGlobPatterns = make([]string, l)
		}
		for k0 := range d.IncludeGlobPatterns {

			{
				l := uint64(0)

				{

					bs := uint8(7)
					t := uint64(buf[i+6] & 0x7F)
					for buf[i+6]&0x80 == 0x80 {
						i++
						t |= uint64(buf[i+6]&0x7F) << bs
						bs += 7
					}
					i++

					l = t

				}
				d.IncludeGlobPatterns[k0] = string(buf[i+6 : i+6+l])
				i += l
			}

		}
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+6] & 0x7F)
			for buf[i+6]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+6]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		if uint64(cap(d.ExcludeGlobPatterns)) >= l {
			d.ExcludeGlobPatterns = d.ExcludeGlobPatterns[:l]
		} else {
			d.ExcludeGlobPatterns = make([]string, l)
		}
		for k0 := range d.ExcludeGlobPatterns {

			{
				l := uint64(0)

				{

					bs := uint8(7)
					t := uint64(buf[i+6] & 0x7F)
					for buf[i+6]&0x80 == 0x80 {
						i++
						t |= uint64(buf[i+6]&0x7F) << bs
						bs += 7
					}
					i++

					l = t

				}
				d.ExcludeGlobPatterns[k0] = string(buf[i+6 : i+6+l])
				i += l
			}

		}
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+6] & 0x7F)
			for buf[i+6]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+6]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		if uint64(cap(d.IncludeRegexPatterns)) >= l {
			d.IncludeRegexPatterns = d.IncludeRegexPatterns[:l]
		} else {
			d.IncludeRegexPatterns = make([]string, l)
		}
		for k0 := range d.IncludeRegexPatterns {

			{
				l := uint64(0)

				{

					bs := uint8(7)
					t := uint64(buf[i+6] & 0x7F)
					for buf[i+6]&0x80 == 0x80 {
						i++
						t |= uint64(buf[i+6]&0x7F) << bs
						bs += 7
					}
					i++

					l = t

				}
				d.IncludeRegexPatterns[k0] = string(buf[i+6 : i+6+l])
				i += l
			}

		}
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+6] & 0x7F)
			for buf[i+6]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+6]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		if uint64(cap(d.ExcludeRegexPatterns)) >= l {
			d.ExcludeRegexPatterns = d.ExcludeRegexPatterns[:l]
		} else {
			d.ExcludeRegexPatterns = make([]string, l)
		}
		for k0 := range d.ExcludeRegexPatterns {

			{
				l := uint64(0)

				{

					bs := uint8(7)
					t := uint64(buf[i+6] & 0x7F)
					for buf[i+6]&0x80 == 0x80 {
						i++
						t |= uint64(buf[i+6]&0x7F) << bs
						bs += 7
					}
					i++

					l = t

				}
				d.ExcludeRegexPatterns[k0] = string(buf[i+6 : i+6+l])
				i += l
			}

		}
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+6] & 0x7F)
			for buf[i+6]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+6]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.UseTitlePrefix = string(buf[i+6 : i+6+l])
		i += l
	}
	{
		d.UseLibraryAsIdentifierPrefix = buf[i+6] == 1
	}
	{
		d.UsePathInLibraryAsIdentifier = buf[i+7] == 1
	}
	{
		d.UseFileNameAsIdentifier = buf[i+8] == 1
	}
	{
		d.UsePathFingerprintAsIdentifier = buf[i+9] == 1
	}
	{
		d.UseFileExtensionAsFormat = buf[i+10] == 1
	}
	return i + 11, nil
}

/*
type IndexGob struct {
	Documents        []*Document
	Libraries        []*Library
	LibraryDocuments []IndexLibraryDocumentsGob
}
*/

func (d *IndexGob) Size() (s uint64) {

	{
		l := uint64(len(d.Documents))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}

		for k0 := range d.Documents {

			{
				if d.Documents[k0] != nil {

					{
						s += (*d.Documents[k0]).Size()
					}
					s += 0
				}
			}

			s += 1

		}

	}
	{
		l := uint64(len(d.Libraries))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}

		for k0 := range d.Libraries {

			{
				if d.Libraries[k0] != nil {

					{
						s += (*d.Libraries[k0]).Size()
					}
					s += 0
				}
			}

			s += 1

		}

	}
	{
		l := uint64(len(d.LibraryDocuments))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}

		for k0 := range d.LibraryDocuments {

			{
				s += d.LibraryDocuments[k0].Size()
			}

		}

	}
	return
}
func (d *IndexGob) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{
		l := uint64(len(d.Documents))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+0] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+0] = byte(t)
			i++

		}
		for k0 := range d.Documents {

			{
				if d.Documents[k0] == nil {
					buf[i+0] = 0
				} else {
					buf[i+0] = 1

					{
						nbuf, err := (*d.Documents[k0]).Marshal(buf[i+1:])
						if err != nil {
							return nil, err
						}
						i += uint64(len(nbuf))
					}
					i += 0
				}
			}

			i += 1

		}
	}
	{
		l := uint64(len(d.Libraries))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+0] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+0] = byte(t)
			i++

		}
		for k0 := range d.Libraries {

			{
				if d.Libraries[k0] == nil {
					buf[i+0] = 0
				} else {
					buf[i+0] = 1

					{
						nbuf, err := (*d.Libraries[k0]).Marshal(buf[i+1:])
						if err != nil {
							return nil, err
						}
						i += uint64(len(nbuf))
					}
					i += 0
				}
			}

			i += 1

		}
	}
	{
		l := uint64(len(d.LibraryDocuments))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+0] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+0] = byte(t)
			i++

		}
		for k0 := range d.LibraryDocuments {

			{
				nbuf, err := d.LibraryDocuments[k0].Marshal(buf[i+0:])
				if err != nil {
					return nil, err
				}
				i += uint64(len(nbuf))
			}

		}
	}
	return buf[:i+0], nil
}

func (d *IndexGob) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i+0]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		if uint64(cap(d.Documents)) >= l {
			d.Documents = d.Documents[:l]
		} else {
			d.Documents = make([]*Document, l)
		}
		for k0 := range d.Documents {

			{
				if buf[i+0] == 1 {
					if d.Documents[k0] == nil {
						d.Documents[k0] = new(Document)
					}

					{
						ni, err := (*d.Documents[k0]).Unmarshal(buf[i+1:])
						if err != nil {
							return 0, err
						}
						i += ni
					}
					i += 0
				} else {
					d.Documents[k0] = nil
				}
			}

			i += 1

		}
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i+0]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		if uint64(cap(d.Libraries)) >= l {
			d.Libraries = d.Libraries[:l]
		} else {
			d.Libraries = make([]*Library, l)
		}
		for k0 := range d.Libraries {

			{
				if buf[i+0] == 1 {
					if d.Libraries[k0] == nil {
						d.Libraries[k0] = new(Library)
					}

					{
						ni, err := (*d.Libraries[k0]).Unmarshal(buf[i+1:])
						if err != nil {
							return 0, err
						}
						i += ni
					}
					i += 0
				} else {
					d.Libraries[k0] = nil
				}
			}

			i += 1

		}
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i+0]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		if uint64(cap(d.LibraryDocuments)) >= l {
			d.LibraryDocuments = d.LibraryDocuments[:l]
		} else {
			d.LibraryDocuments = make([]IndexLibraryDocumentsGob, l)
		}
		for k0 := range d.LibraryDocuments {

			{
				ni, err := d.LibraryDocuments[k0].Unmarshal(buf[i+0:])
				if err != nil {
					return 0, err
				}
				i += ni
			}

		}
	}
	return i + 0, nil
}

/*
type IndexLibraryDocumentsGob struct {
	Library   string
	Documents []string
}
*/

func (d *IndexLibraryDocumentsGob) Size() (s uint64) {

	{
		l := uint64(len(d.Library))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}
		s += l
	}
	{
		l := uint64(len(d.Documents))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}

		for k0 := range d.Documents {

			{
				l := uint64(len(d.Documents[k0]))

				{

					t := l
					for t >= 0x80 {
						t >>= 7
						s++
					}
					s++

				}
				s += l
			}

		}

	}
	return
}
func (d *IndexLibraryDocumentsGob) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{
		l := uint64(len(d.Library))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+0] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+0] = byte(t)
			i++

		}
		copy(buf[i+0:], d.Library)
		i += l
	}
	{
		l := uint64(len(d.Documents))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+0] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+0] = byte(t)
			i++

		}
		for k0 := range d.Documents {

			{
				l := uint64(len(d.Documents[k0]))

				{

					t := uint64(l)

					for t >= 0x80 {
						buf[i+0] = byte(t) | 0x80
						t >>= 7
						i++
					}
					buf[i+0] = byte(t)
					i++

				}
				copy(buf[i+0:], d.Documents[k0])
				i += l
			}

		}
	}
	return buf[:i+0], nil
}

func (d *IndexLibraryDocumentsGob) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i+0]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		d.Library = string(buf[i+0 : i+0+l])
		i += l
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i+0]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		if uint64(cap(d.Documents)) >= l {
			d.Documents = d.Documents[:l]
		} else {
			d.Documents = make([]string, l)
		}
		for k0 := range d.Documents {

			{
				l := uint64(0)

				{

					bs := uint8(7)
					t := uint64(buf[i+0] & 0x7F)
					for buf[i+0]&0x80 == 0x80 {
						i++
						t |= uint64(buf[i+0]&0x7F) << bs
						bs += 7
					}
					i++

					l = t

				}
				d.Documents[k0] = string(buf[i+0 : i+0+l])
				i += l
			}

		}
	}
	return i + 0, nil
}
