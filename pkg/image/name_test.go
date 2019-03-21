/*
 * Copyright 2019 The original author or authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package image_test

import (
	"cnab-k8s-installer-base/pkg/image"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Name", func() {

	var (
		ref image.Name
	)

	Describe("NewName", func() {
		var (
			name string
			err  error
		)

		JustBeforeEach(func() {
			ref, err = image.NewName(name)
		})

		Context("when the string name is empty", func() {
			BeforeEach(func() {
				name = ""
			})

			It("should return a suitable error", func() {
				Expect(err).To(MatchError("invalid reference format"))
			})
		})

		Context("when the string name contains no tag or digest", func() {
			BeforeEach(func() {
				name = "ubuntu"
			})

			It("should succeed", func() {
				Expect(err).NotTo(HaveOccurred())
			})

			It("should not include a tag", func() {
				Expect(ref.Tag()).To(Equal(""))
			})

			It("should not include a digest", func() {
				Expect(ref.Digest()).To(Equal(image.EmptyDigest))
			})

			It("should return a suitable string form", func() {
				Expect(ref.String()).To(Equal("docker.io/library/ubuntu"))
			})

			It("should return the correct path", func() {
				Expect(ref.Path()).To(Equal("library/ubuntu"))
			})

			It("should normalize to itself", func() {
				Expect(ref.Normalize()).To(Equal(ref))
			})

			It("should return the correct synonyms", func() {
				Expect(synonymStrings(ref)).To(ConsistOf("ubuntu", "library/ubuntu", "docker.io/library/ubuntu", "index.docker.io/library/ubuntu"))
			})

			It("should return the correct name", func() {
				Expect(ref.Name()).To(Equal("docker.io/library/ubuntu"))
			})
		})

		Context("when the string name includes a tag", func() {
			BeforeEach(func() {
				name = "ubuntu:18.10"
			})

			It("should succeed", func() {
				Expect(err).NotTo(HaveOccurred())
			})

			It("should include the tag", func() {
				Expect(ref.Tag()).To(Equal("18.10"))
			})

			It("should return a suitable string form", func() {
				Expect(ref.String()).To(Equal("docker.io/library/ubuntu:18.10"))
			})

			It("should return the correct path", func() {
				Expect(ref.Path()).To(Equal("library/ubuntu"))
			})

			It("should normalize to itself", func() {
				Expect(ref.Normalize()).To(Equal(ref))
			})

			It("should return the correct synonyms", func() {
				Expect(synonymStrings(ref)).To(ConsistOf("ubuntu:18.10", "library/ubuntu:18.10", "docker.io/library/ubuntu:18.10", "index.docker.io/library/ubuntu:18.10"))
			})

			It("should return the correct name", func() {
				Expect(ref.Name()).To(Equal("docker.io/library/ubuntu"))
			})
		})

		Context("when the string name includes a digest", func() {
			BeforeEach(func() {
				name = "ubuntu@sha256:deadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef"
			})

			It("should succeed", func() {
				Expect(err).NotTo(HaveOccurred())
			})

			It("should include the digest", func() {
				Expect(ref.Digest()).To(Equal(image.NewDigest("sha256:deadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef")))
			})

			It("should return a suitable string form", func() {
				Expect(ref.String()).To(Equal("docker.io/library/ubuntu@sha256:deadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef"))
			})

			It("should return the correct path", func() {
				Expect(ref.Path()).To(Equal("library/ubuntu"))
			})

			It("should normalize to itself", func() {
				Expect(ref.Normalize()).To(Equal(ref))
			})

			It("should return the correct synonyms", func() {
				Expect(synonymStrings(ref)).To(ConsistOf("ubuntu@sha256:deadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef",
					"library/ubuntu@sha256:deadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef",
					"docker.io/library/ubuntu@sha256:deadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef",
					"index.docker.io/library/ubuntu@sha256:deadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef"))
			})

			It("should return the correct name", func() {
				Expect(ref.Name()).To(Equal("docker.io/library/ubuntu"))
			})
		})

		Context("when the string name includes a tag and a digest", func() {
			BeforeEach(func() {
				name = "ubuntu:18.10@sha256:deadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef"
			})

			It("should succeed", func() {
				Expect(err).NotTo(HaveOccurred())
			})

			It("should include the tag", func() {
				Expect(ref.Tag()).To(Equal("18.10"))
			})

			It("should include the digest", func() {
				Expect(ref.Digest()).To(Equal(image.NewDigest("sha256:deadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef")))
			})

			It("should return a suitable string form", func() {
				Expect(ref.String()).To(Equal("docker.io/library/ubuntu:18.10@sha256:deadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef"))
			})

			It("should return the correct path", func() {
				Expect(ref.Path()).To(Equal("library/ubuntu"))
			})

			It("should normalize to itself", func() {
				Expect(ref.Normalize()).To(Equal(ref))
			})

			It("should return the correct synonyms", func() {
				Expect(synonymStrings(ref)).To(ConsistOf("ubuntu:18.10@sha256:deadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef",
					"library/ubuntu:18.10@sha256:deadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef",
					"docker.io/library/ubuntu:18.10@sha256:deadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef",
					"index.docker.io/library/ubuntu:18.10@sha256:deadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef"))
			})

			It("should return the correct name", func() {
				Expect(ref.Name()).To(Equal("docker.io/library/ubuntu"))
			})
		})

		Describe("synonyms", func() {
			Context("when the string name contains library", func() {
				BeforeEach(func() {
					name = "library/ubuntu"
				})

				It("should succeed", func() {
					Expect(err).NotTo(HaveOccurred())
				})

				It("should return a suitable string form", func() {
					Expect(ref.String()).To(Equal("docker.io/library/ubuntu"))
				})

				It("should return the correct synonyms", func() {
					Expect(synonymStrings(ref)).To(ConsistOf("ubuntu", "library/ubuntu", "docker.io/library/ubuntu", "index.docker.io/library/ubuntu"))
				})
			})

			Context("when the string name contains the hostname docker.io", func() {
				BeforeEach(func() {
					name = "docker.io/library/ubuntu"
				})

				It("should succeed", func() {
					Expect(err).NotTo(HaveOccurred())
				})

				It("should return a suitable string form", func() {
					Expect(ref.String()).To(Equal("docker.io/library/ubuntu"))
				})

				It("should return the correct synonyms", func() {
					Expect(synonymStrings(ref)).To(ConsistOf("ubuntu", "library/ubuntu", "docker.io/library/ubuntu", "index.docker.io/library/ubuntu"))
				})
			})

			Context("when the string name contains the hostname index.docker.io", func() {
				BeforeEach(func() {
					name = "index.docker.io/library/ubuntu"
				})

				It("should succeed", func() {
					Expect(err).NotTo(HaveOccurred())
				})

				It("should return a suitable string form", func() {
					Expect(ref.String()).To(Equal("docker.io/library/ubuntu"))
				})

				It("should return the correct synonyms", func() {
					Expect(synonymStrings(ref)).To(ConsistOf("ubuntu", "library/ubuntu", "docker.io/library/ubuntu", "index.docker.io/library/ubuntu"))
				})
			})

			Describe("synonyms of synonyms", func() {
				BeforeEach(func() {
					name = "index.docker.io/library/ubuntu"
				})

				It("should not produce new synonyms", func() {
					synonyms := ref.Synonyms()
					for _, s := range synonyms {
						// Normalize to ensure synonyms can be computed
						Expect(s.Normalize().Synonyms()).To(ConsistOf(synonyms))
					}
				})
			})

			Describe("behaviour of synonyms which do not include a host name", func() {
				BeforeEach(func() {
					name = "ubuntu:18.10@sha256:deadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef"
				})

				It("should behave correctly", func() {
					// A Name without a host name can only be generated using the Synonyms method.
					synonyms := ref.Synonyms()
					for _, s := range synonyms {
						// If the synonym does not have a host name, check its behaviour.
						if s.String() == name {
							By("Tag")
							Expect(s.Tag()).To(Equal("18.10"))

							By("Digest")
							Expect(s.Digest()).To(Equal(image.NewDigest("sha256:deadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef")))

							By("Path")
							Expect(s.Path()).To(Equal("library/ubuntu"))

							By("Normalize")
							Expect(s.Normalize()).To(Equal(ref))

							By("Synonyms")
							Expect(synonymStrings(s)).To(ConsistOf("ubuntu:18.10@sha256:deadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef",
								"library/ubuntu:18.10@sha256:deadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef",
								"docker.io/library/ubuntu:18.10@sha256:deadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef",
								"index.docker.io/library/ubuntu:18.10@sha256:deadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef"))
						}
					}
				})
			})
		})
	})

	Describe("EmptyName", func() {
		JustBeforeEach(func() {
			ref = image.EmptyName
		})

		It("should return an empty digest", func() {
			Expect(ref.Digest()).To(Equal(image.EmptyDigest))
		})

		It("should return an empty tag", func() {
			Expect(ref.Tag()).To(Equal(""))
		})

		It("should return an empty string form", func() {
			Expect(ref.String()).To(Equal(""))
		})

		It("should return itself as the only synonym", func() {
			Expect(ref.Synonyms()).To(ConsistOf(image.EmptyName))
		})

		It("should normalize to itself", func() {
			Expect(ref.Normalize()).To(Equal(ref))
		})

		It("should panic when asked for its path", func() {
			Expect(func() { ref.Path() }).To(Panic())
		})
	})

	Describe("WithTag", func() {
		var (
			newRef image.Name
			tag    string
			err    error
		)

		JustBeforeEach(func() {
			newRef, err = ref.WithTag(tag)
		})

		Context("when the tag is valid", func() {
			BeforeEach(func() {
				tag = "test-tag"
			})

			Context("when the image name is tagged", func() {
				BeforeEach(func() {
					var err error
					ref, err = image.NewName("ubuntu:some-tag")
					Expect(err).NotTo(HaveOccurred())
				})

				It("should replace the tag", func() {
					Expect(err).NotTo(HaveOccurred())
					Expect(newRef.Tag()).To(Equal("test-tag"))
				})
			})

			Context("when the image name is not tagged", func() {
				BeforeEach(func() {
					var err error
					ref, err = image.NewName("ubuntu")
					Expect(err).NotTo(HaveOccurred())
				})

				It("should set the tag", func() {
					Expect(err).NotTo(HaveOccurred())
					Expect(newRef.Tag()).To(Equal("test-tag"))
					Expect(newRef.Path()).To(Equal("library/ubuntu"))
				})
			})
		})

		Context("when the tag is invalid", func() {
			BeforeEach(func() {
				tag = "-invalid"
				var err error
				ref, err = image.NewName("ubuntu")
				Expect(err).NotTo(HaveOccurred())
			})

			It("should return a suitable error", func() {
				Expect(err).To(MatchError("Cannot apply tag -invalid to image.Name docker.io/library/ubuntu: invalid tag format"))
			})
		})
	})

	Describe("WithoutTag", func() {
		var newRef image.Name

		JustBeforeEach(func() {
			newRef = ref.WithoutTag()
		})

		Context("when the image name is tagged", func() {
			BeforeEach(func() {
				var err error
				ref, err = image.NewName("ubuntu:some-tag")
				Expect(err).NotTo(HaveOccurred())
			})

			It("should remove the tag", func() {
				Expect(newRef.Tag()).To(Equal(""))
				Expect(newRef.String()).To(Equal("docker.io/library/ubuntu"))
			})
		})

		Context("when the image name is not tagged", func() {
			BeforeEach(func() {
				var err error
				ref, err = image.NewName("ubuntu")
				Expect(err).NotTo(HaveOccurred())
			})

			It("should not change the image name", func() {
				Expect(newRef).To(Equal(ref))
			})
		})
	})

	Describe("WithDigest", func() {
		var (
			newRef image.Name
			digest image.Digest
			err    error
		)

		JustBeforeEach(func() {
			newRef, err = ref.WithDigest(digest)
		})

		Context("when the digest is valid", func() {
			BeforeEach(func() {
				digest = image.NewDigest("sha256:2fb7bfc6145d0ad40334f1802707c2e2390bdcfc16ca636d9ed8a56c1101f5b9")
			})

			Context("when the image name is tagged", func() {
				BeforeEach(func() {
					var err error
					ref, err = image.NewName("ubuntu:some-tag")
					Expect(err).NotTo(HaveOccurred())
				})

				It("should keep the tag and set the digest", func() {
					Expect(err).NotTo(HaveOccurred())
					Expect(newRef.Tag()).To(Equal("some-tag"))
					Expect(newRef.Digest().String()).To(Equal("sha256:2fb7bfc6145d0ad40334f1802707c2e2390bdcfc16ca636d9ed8a56c1101f5b9"))
					Expect(newRef.String()).To(Equal("docker.io/library/ubuntu:some-tag@sha256:2fb7bfc6145d0ad40334f1802707c2e2390bdcfc16ca636d9ed8a56c1101f5b9"))
				})
			})

			Context("when the image name is not tagged", func() {
				BeforeEach(func() {
					var err error
					ref, err = image.NewName("ubuntu")
					Expect(err).NotTo(HaveOccurred())
				})

				It("should set the digest", func() {
					Expect(err).NotTo(HaveOccurred())
					Expect(newRef.Tag()).To(Equal(""))
					Expect(newRef.Digest().String()).To(Equal("sha256:2fb7bfc6145d0ad40334f1802707c2e2390bdcfc16ca636d9ed8a56c1101f5b9"))
					Expect(newRef.String()).To(Equal("docker.io/library/ubuntu@sha256:2fb7bfc6145d0ad40334f1802707c2e2390bdcfc16ca636d9ed8a56c1101f5b9"))
				})
			})
		})

		Context("when the digest is invalid", func() {
			BeforeEach(func() {
				digest = image.NewDigest("2fb7bfc6145d0ad40334f1802707c2e2390bdcfc16ca636d9ed8a56c1101f5b9")
				var err error
				ref, err = image.NewName("ubuntu")
				Expect(err).NotTo(HaveOccurred())
			})

			It("should return a suitable error", func() {
				Expect(err).To(MatchError("Cannot apply digest 2fb7bfc6145d0ad40334f1802707c2e2390bdcfc16ca636d9ed8a56c1101f5b9 to image.Name docker.io/library/ubuntu: invalid digest format"))
			})
		})
	})
})

func synonymStrings(ref image.Name) []string {
	ss := []string{}
	for _, s := range ref.Synonyms() {
		ss = append(ss, s.String())
	}
	return ss
}
