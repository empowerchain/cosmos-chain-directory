package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/fs"
	"net/http"
)

//go:embed chain-registry
var chainRegistry embed.FS

type AllChainsResponse struct {
	Chains []string `json:"chains"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", AllChainsHandler).Methods("GET")
	r.HandleFunc("/{chain}", ChainHandler).Methods("GET")

	fmt.Println("Serving on port 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}

func AllChainsHandler(w http.ResponseWriter, _ *http.Request) {
	var chains []string

	dirEntries, err := fs.ReadDir(chainRegistry, "chain-registry")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for _, e := range dirEntries {
		if e.IsDir() {
			chains = append(chains, e.Name())
		}
	}

	res := AllChainsResponse{
		Chains: chains,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func ChainHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	chain := vars["chain"]

	path := fmt.Sprintf("chain-registry/%s/chain.json", chain)
	chainInfo, err := fs.ReadFile(chainRegistry, path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(chainInfo)
}
