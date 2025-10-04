package app 

import "net/http"

type APP struct { 
	router http.Handler
} 
func New() *APP{ 
 return &APP{ 
   reouter chi.NewRouter()
	} 
} 





