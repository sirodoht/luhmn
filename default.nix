with (import <nixpkgs> {});
mkShell {
  buildInputs = [
    go
    postgresql_12
  ];
  shellHook = ''
    export GOPATH="$(pwd)/go";
  '';
}
