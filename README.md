# mantela-adder

Mantela Adder enables you to add mantela files to your mantela collection.

## Usage

1. copy `config.sample.toml` to `config.toml` and change it as you like.

2. save mantela file that you want to override to source mantela file.

3. run mantela-adder.

## Tips

if the destination of `diff` is not set in `config.toml`, mantela-adder serves `source` mantela simply.

If mantela.json is fetched from the web, it can be treated as a proxy, and if mantela.json is served from a file, it can be treated as mantela.json hosting software.

This means that it can be treated as a mantela.json server that can be easily hosted in a virtual environment.
