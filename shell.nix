with (import (fetchTarball
  "https://github.com/nixos/nixpkgs/archive/9e53905d6bb02134e86117d526ee62324047e863.tar.gz")
  { });

mkShell { buildInputs = [ go_1_19 docker usql ]; }
