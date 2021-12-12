# Generate module graphs


## Get modgv

go install github.com/lucasepe/modgv/modgv


go mod graph | modgv

Generates graphviz dot language from module graph


## Graphviz

apk update
apk add graphviz

go mod graph | modgv | dot -Tpng -o graph.png
