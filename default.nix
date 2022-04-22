with (import <nixpkgs> {});
mkShell {
  buildInputs = [
    postgresql_12
    go_1_17
  ];
  shellHook = ''
    export GOPATH="$(pwd)/go";
  '';
}
