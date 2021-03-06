package image

import (
	"fmt"
	"path"
	"strings"

	"github.com/docker/distribution/reference"
)

const (
	dockerHubHost     = "docker.io"
	fullDockerHubHost = "index.docker.io"
)

// Name is a named image reference
type Name struct {
	ref reference.Named
}

var EmptyName Name

func init() {
	EmptyName = Name{nil}
}

func NewName(i string) (Name, error) {
	ref, err := reference.ParseNormalizedNamed(i)
	return Name{ref}, err
}

func (img Name) Normalize() Name {
	if img.ref == nil {
		return EmptyName
	}
	ref, err := NewName(img.String())
	if err != nil {
		panic(err) // should never happen
	}
	return ref
}

func (img Name) Name() string {
	return img.ref.Name()
}

func (img Name) String() string {
	if img.ref == nil {
		return ""
	}
	return img.ref.String()
}

func (img Name) Host() string {
	h, _ := img.parseHostPath()
	return h
}

func (img Name) Path() string {
	_, p := img.parseHostPath()
	return p
}

func (img Name) Tag() string {
	if taggedRef, ok := img.ref.(reference.Tagged); ok {
		return taggedRef.Tag()
	}
	return ""
}

func (img Name) WithTag(tag string) (Name, error) {
	namedTagged, err := reference.WithTag(img.ref, tag)
	if err != nil {
		return EmptyName, fmt.Errorf("Cannot apply tag %s to image.Name %v: %v", tag, img, err)
	}
	return Name{namedTagged}, nil
}

func (img Name) Digest() Digest {
	if digestedRef, ok := img.ref.(reference.Digested); ok {
		return NewDigest(string(digestedRef.Digest()))
	}
	return EmptyDigest
}

func (img Name) WithoutTag() Name {
	return Name{reference.TrimNamed(img)}
}

func (img Name) WithDigest(digest Digest) (Name, error) {
	digested, err := reference.WithDigest(img.ref, digest.dig)
	if err != nil {
		return EmptyName, fmt.Errorf("Cannot apply digest %s to image.Name %v: %v", digest, img, err)
	}

	return Name{digested}, nil
}

func (img Name) WithoutDigest() string {
	if strings.Contains(img.ref.Name(), "@") {
		return strings.Split(img.ref.Name(), "@")[0]
	}
	return img.ref.Name()
}

// Synonyms returns the equivalent image names for a given image name. The synonyms are not necessarily
// normalized: in particular they may not have a host name.
func (img Name) Synonyms() []Name {
	if img.ref == nil {
		return []Name{EmptyName}
	}
	imgHost, imgRepoPath := img.parseHostPath()
	nameMap := map[Name]struct{}{img: {}}

	if imgHost == dockerHubHost {
		elidedImg := imgRepoPath
		name, err := synonym(img, elidedImg)
		if err == nil {
			nameMap[name] = struct{}{}
		}

		elidedImgElements := strings.Split(elidedImg, "/")
		if len(elidedImgElements) == 2 && elidedImgElements[0] == "library" {
			name, err := synonym(img, elidedImgElements[1])
			if err == nil {
				nameMap[name] = struct{}{}
			}
		}

		fullImg := path.Join(fullDockerHubHost, imgRepoPath)
		name, err = synonym(img, fullImg)
		if err == nil {
			nameMap[name] = struct{}{}
		}

		dockerImg := path.Join(dockerHubHost, imgRepoPath)
		name, err = synonym(img, dockerImg)
		if err == nil {
			nameMap[name] = struct{}{}
		}
	}

	names := []Name{}
	for n, _ := range nameMap {
		names = append(names, n)
	}

	return names
}

func synonym(original Name, newName string) (Name, error) {
	named, err := reference.WithName(newName)
	if err != nil {
		return EmptyName, err
	}

	if taggedRef, ok := original.ref.(reference.Tagged); ok {
		named, err = reference.WithTag(named, taggedRef.Tag())
		if err != nil {
			return EmptyName, err
		}
	}

	if digestedRef, ok := original.ref.(reference.Digested); ok {
		named, err = reference.WithDigest(named, digestedRef.Digest())
		if err != nil {
			return EmptyName, err
		}
	}

	return Name{named}, nil
}

func (img Name) parseHostPath() (host string, repoPath string) {
	s := strings.SplitN(img.ref.Name(), "/", 2)
	if len(s) == 1 {
		return img.Normalize().parseHostPath()
	}
	return s[0], s[1]
}
