# dvd

`dvd` is a command line utility to divide a list of directories into smaller chunks, to run your tests in parallel.

## A brief background

Let's say you have the following Go modules.

    $ ls lib
    cat
    dog
    sheep

The `lib` directory is small enough to complete your tests in a few seconds.
But consider a case when a bunch of new modules comes into the `lib` directory.

    $ ls lib
    alpaca
    ant
    baboon
    bat
    bird
    camel
    cat
    cow
    dog
    duck

More modules slows your test completion, and you'll soon lose your focus and productivity.

<div style="text-align: center">
  <a href="https://xkcd.com/303/">
    <image src="https://imgs.xkcd.com/comics/compiling.png">
  </a>
</div>

This is where `dvd` comes in.
When refactoring your tests to run in parallel, give the directory, machine sequence and max parallelism to `dvd`,
to divide long list of modules into reasonable chunks.

Suppose you have 3 machines to run in parallel, set max parallelism to 3 and give the machine number as a sequence.
`dvd` will evenly distribute the chunks so that each test chunks completes in reasonable duration.

    $ dvd -dir lib -sequence 0 -parallel 3
    alpaca
    ant
    baboon

    $ dvd -dir lib -sequence 1 -parallel 3
    bat
    bird
    camel

    $ dvd -dir lib -sequence 2 -parallel 3
    cat
    cow
    dog
    duck

## GitHub Actions example

Easily distribute your test runs with [strategy.martix](https://docs.github.com/en/free-pro-team@latest/actions/reference/workflow-syntax-for-github-actions#jobsjob_idstrategymatrix) property.

    strategy:
      matrix:
        include:
          - dir: path/to/lib
            sequence: 0
            parallel: 3

          - dir: path/to/lib
            sequence: 1
            parallel: 3

          - dir: path/to/lib
            sequence: 2
            parallel: 3
