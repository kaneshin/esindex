# Elasticsearch Index

## Installation

```
go get github.com/kaneshin/esindex
```


## Usage

### Create

Create an index, appending a timestamp to the given name. And create an alias which points to the index if it isn't associated.

```
esindex create my_index --url http://127.0.0.1:9200 --mappings '{...mappings...}'
```

```
curl -XPUT http://127.0.0.1:9200/my_index_20060102150405 -d '{...mappings...}'
curl -XPOST http://127.0.0.1:9200/_aliases -d '
{
    "actions": [{
        "add": {
            "alias": "my_index",
            "index": "my_index_20060102150405"
        }}
    ]
}
'
```

### List

Show the my_index aliases.

```
esindex list my_index --url http://127.0.0.1:9200
```


### Alias

Change the my_index alias to point to the given name.

```
esindex alias my_index_20091110230000 --url http://127.0.0.1:9200
```

```
curl -XPOST http://127.0.0.1:9200/_aliases -d '
{
    "actions": [{
        "remove": {
            "alias": "my_index",
            "index": "my_index_20060102150405"
        }
    }, {
        "add": {
            "alias": "my_index",
            "index": "my_index_20091110230000"
        }
    }]
}
'
```


## Tutorial

```
$ esindex create my_index --url http://127.0.0.1:9200 --mappings '
{
    "mappings": {
        "articles": {
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    }
}
'
my_index_20151024000500 <- my_index


$ esindex create my_index --url http://127.0.0.1:9200 --mappings '
{
    "mappings": {
        "articles": {
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "rating": {
                    "type": "integer"
                }
            }
        }
    }
}
'
my_index_20151025111000


$ esindex list my_index --url http://127.0.0.1:9200
my_index_20151024000500 <- my_index
my_index_20151025111000


$ esindex alias my_index_20151025111000 --url http://127.0.0.1:9200
my_index_20151025111000 <- my_index


$ esindex list my_index --url http://127.0.0.1:9200
my_index_20151024000500
my_index_20151025111000 <- my_index
```
