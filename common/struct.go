/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2019-01-12
 * Time: 19:43
 */
package common

type PostStore struct {
	Title string `json:"title"`
	Category int `json:"category"`
	Tags []int `json:"tags"`
	Summary string `json:"summary"`
	Content string `json:"content"`
}