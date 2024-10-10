# Markdown manual test documentation


<img src="img/gopher.jpg" alt="drawing" width="200"/>

The Task:

To document manual a **test case** you have to document

- the test case, which includes
  - Needed steps
  - Expected results

To document a single **test run** you have to copy the test ‚case and fill in the results.

Or you can by an expensive licence.

## Solution

1) Document the test case in a markdown file "testcase1.md"
2) For a test run, only change:
    - `- [ ]` to `- [x]`
    - Fill in results after `>:`
3) After the test, call

```bash
task render
```

Output:

```log
task render
task: [render] markmac <testcase1.md > tmp.md
task: [render] pandoc tmp.md -o result-testcase1-dev-1.0.0-2024-10-10-1842.pdf  -V papersize:a4 -V geometry:margin=1in
task: [render] open result-testcase1-dev-1.0.0-2024-10-10-1842.pdf
```

You get an pdf with the result, see [PDF Results](result-testcase1-dev-1.0.0-2024-10-10-1842.pdf)

4) To reset any entries, run

```bash
task refresh
```

All checkboxes will be reset and all text behind an `>:` will be replaces with `___`.

## Installation

- Install task.dev
- Install go
- run `task install`
