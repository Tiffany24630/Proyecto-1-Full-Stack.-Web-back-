package main

type Series struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Type     string `json:"type"`
	Total    int    `json:"total"`
	Progress int    `json:"progress"`
	Image    string `json:"image"`
}
