# jsoninf

JSON schema inference, especially from JSON lines.

* stage: "mvp", "wip"

# Problem

Given a blob of JSON documents w/o explicit schema, infer structure and data
types. This tool can be used to detect structural and typing issue in large
JSON datasets.

# Usage

```shell
$ jsoninf < file.ndj
```

# Examples

```shell
$ cat fixtures/a.json
{"a": "hello"}
{"a": "world"}

$ jsoninf < fixtures/a.json
/a [string]

$ cat fixtures/b.json
{"a": "hello"}
{"a": 123 }

$ jsoninf < fixtures/b.json
2021/02/13 01:47:47 line 2: mixed types detected in: /a [string, float64]
/a [string]
```

To only print errors:

```shell
$ jsoninf < fixtures/b.json > /dev/null
2021/02/13 01:47:47 line 2: mixed types detected in: /a [string, float64]
```

Example:

```
$ jsoninf < fixtures/ref.json
/biblio/arxiv_id [string]
/biblio/container_name [string]
/biblio/contrib_raw_names [slice]
/biblio/doi [string]
/biblio/issue [string]
/biblio/pages [string]
/biblio/pmcid [string]
/biblio/publisher [string]
/biblio/title [string]
/biblio/unstructured [string]
/biblio/url [string]
/biblio/volume [string]
/biblio/year [float64]
/index [float64]
/key [string]
/ref_source [string]
/release_ident [string]
/release_year [float64]
/work_ident [string]
2021/02/13 02:21:47 found 0 issues in 1 lines
```

# Performance

Disclaimer: This is totally unoptimized code.

* checking 1M lines takes about 7min

# Literature

* [Schema Inference for Massive JSON Datasets](https://openproceedings.org/2017/conf/edbt/paper-62.pdf)
* [Human-in-the-Loop Schema Inference for Massive JSONDatasets](https://openproceedings.org/2020/conf/edbt/paper_318.pdf)

