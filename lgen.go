package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)



func handler(w http.ResponseWriter, r *http.Request) {
	rpg := rand.New(rand.NewSource(time.Now().UnixNano()))

	fmt.Fprintf(w, "%s\n", "# TYPE metric_without_timestamp_and_labels gauge")
	fmt.Fprintf(w, "metric_without_timestamp_and_labels %3.1f\n", rpg.Float32() * 100)

	fmt.Fprintf(w, "%s\n","# TYPE metric_without_timestamp_and_labels2 gauge")
	fmt.Fprintf(w, "metric_without_timestamp_and_labels2 %3.1f\n", rpg.Float32() * 42)

	fmt.Fprintf(w, "%s\n","# TYPE metric_with_label counter")
	fmt.Fprintf(w, "metric_with_label{code=\"200\"} %3.0f\n", rpg.Float32() * 42)
	fmt.Fprintf(w, "metric_with_label{code=\"500\"} %3.0f\n", rpg.Float32() * 42)



	time.Sleep( time.Duration(rpg.Float32() * 40 )* time.Millisecond)
}

func main() {
	http.HandleFunc("/metrics", handler)
	http.ListenAndServe(":8080", nil)
}
