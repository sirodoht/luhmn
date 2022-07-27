with (import <nixpkgs> {});
mkShell {
  buildInputs = [
    postgresql_12
    go
    golangci-lint
  ];
  shellHook = ''
    export GOPATH="$(pwd)/go";
  '';
}
