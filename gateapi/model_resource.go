/*
 * Spinnaker API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"os"
)

type Resource struct {
	Description string `json:"description,omitempty"`
	Filename string `json:"filename,omitempty"`
	Open bool `json:"open,omitempty"`
	Url *Url `json:"url,omitempty"`
	Uri *Uri `json:"uri,omitempty"`
	InputStream *InputStream `json:"inputStream,omitempty"`
	File **os.File `json:"file,omitempty"`
	Readable bool `json:"readable,omitempty"`
}