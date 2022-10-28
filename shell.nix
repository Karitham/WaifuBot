{ pkgs ? import (fetchTarball
  "https://github.com/NixOS/nixpkgs/archive/4fce8949409b4eb6250edf612cf30ab9a94c0da6.tar.gz")
  { } }:

pkgs.mkShell { buildInputs = with pkgs; [ go_1_19 docker usql dbmate sqlc ]; }
