/*
 * VMC Migration Hub - VMware Docs Article Search API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ArticleSearchRequestQueryString struct {

	And string `json:"and,omitempty"`

	MustNot string `json:"mustNot,omitempty"`

	Or string `json:"or,omitempty"`

	Phrase string `json:"phrase,omitempty"`
}