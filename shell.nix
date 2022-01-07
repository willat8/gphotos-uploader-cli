with import (fetchTarball https://github.com/NixOS/nixpkgs-channels/archive/7e9b0dff974c89e070da1ad85713ff3c20b0ca97.tar.gz) {}; {
  devEnv = stdenv.mkDerivation {
    name = "dev";
    buildInputs = [ stdenv go_1_14 glibc.static ];
  };
}
