# Data manifest

Provenance for the files in `data/jolpica/`. The CSVs themselves are not committed (see `.gitignore`); this file records exactly which dump the code was written and tested against.

## Source

| | |
| --- | --- |
| Dataset | Jolpica-F1 database dump (successor to the Ergast API) |
| Download URL | <https://api.jolpi.ca/data/dumps/download/delayed/?dump_type=csv> |
| Dump type | `delayed` CSV: the most recent races may be missing |
| Downloaded | 2026-10-02 |
| Race results through | 2026-09-13 (latest race with results; later 2026 rounds have no entries) |
| Licence | [CC BY-NC-SA 4.0](https://creativecommons.org/licenses/by-nc-sa/4.0/), per the [Jolpica terms of use](https://github.com/jolpica/jolpica-f1/blob/main/TERMS.md): non-commercial use, attribution required, share-alike |

The dump ships every table with a `formula_one_` prefix. Only the 9 tables below were copied, with the prefix removed. Lap times and pit stops were left out.

## Files

Every file is UTF-8 with no BOM, CRLF line endings, and one header row.

| File | Rows (excl. header) | SHA-256 |
| --- | ---: | --- |
| `circuit.csv` | 78 | `df17d7c6dac3c00dbba3f0cbe6942e9f939feb9d1f37e691d40cef85776f821a` |
| `driver.csv` | 881 | `162411f8f32dc17cbbdef380b3a4cdd7c32327eb563c6041f9a3a6e265edfa37` |
| `round.csv` | 1,173 | `1a843144e851e61ef01f9164129f0ec48dd5802092ac74404107560e353c6154` |
| `roundentry.csv` | 27,619 | `f090908a1af6a86efd3f7ac322311900a205c51376a13a19b25ff892c833f5fd` |
| `season.csv` | 77 | `8105785c76e4e8bc87198b22ba7210de56a05dc633ce0f754ba6c2a67b594c12` |
| `session.csv` | 5,130 | `f2dab264640c8f81797f11217d1a044e147be72a169c58a69ddd1aef62365631` |
| `sessionentry.csv` | 51,413 | `0ecd94e3da6bcd554dc5c3e99abd3fc7f28606eb9e8c71493bf2314e9130d1a0` |
| `team.csv` | 214 | `1ba1e7a567d6f326d249a07d15b861ecc15a0b0ad465d9790b47959adc7c1b9a` |
| `teamdriver.csv` | 3,630 | `aaff3853ad5ee5a532c4b94e34e3a701efd35984935406f9a89b1afd193a677c` |

## Verify your copy

From the repo root:

```sh
cd data/jolpica && shasum -a 256 -c <<'EOF'
df17d7c6dac3c00dbba3f0cbe6942e9f939feb9d1f37e691d40cef85776f821a  circuit.csv
162411f8f32dc17cbbdef380b3a4cdd7c32327eb563c6041f9a3a6e265edfa37  driver.csv
1a843144e851e61ef01f9164129f0ec48dd5802092ac74404107560e353c6154  round.csv
f090908a1af6a86efd3f7ac322311900a205c51376a13a19b25ff892c833f5fd  roundentry.csv
8105785c76e4e8bc87198b22ba7210de56a05dc633ce0f754ba6c2a67b594c12  season.csv
f2dab264640c8f81797f11217d1a044e147be72a169c58a69ddd1aef62365631  session.csv
0ecd94e3da6bcd554dc5c3e99abd3fc7f28606eb9e8c71493bf2314e9130d1a0  sessionentry.csv
1ba1e7a567d6f326d249a07d15b861ecc15a0b0ad465d9790b47959adc7c1b9a  team.csv
aaff3853ad5ee5a532c4b94e34e3a701efd35984935406f9a89b1afd193a677c  teamdriver.csv
EOF
```

Every line should print `OK`. If you download a newer dump, put it in a new folder, run the loader's checks on it, then update this manifest in the same commit as any code that depends on it.
