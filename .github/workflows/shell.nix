{ pkgs ? import <nixpkgs> {} }:
let
  python-with-my-packages = pkgs.python3.withPackages (p: with p; [
    pyyaml
    jinja2
  ]);
in
python-with-my-packages.env