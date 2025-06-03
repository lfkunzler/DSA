{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-25.05";
    systems.url = "github:nix-systems/default";
    devenv.url = "github:cachix/devenv/python-rewrite";
  };

  outputs = { nixpkgs, systems, devenv, ... }@inputs:
    let forEachSystem = nixpkgs.lib.genAttrs (import systems);
    in {
      devShells = forEachSystem (system:
        let
          pkgs = nixpkgs.legacyPackages.${system};

          codeGenPackages = [ pkgs.git ];

          languagePkgs = [ pkgs.go ];

          devPackages = with pkgs; [
            # Formatters.
            gofumpt
            hclfmt
            nixfmt-classic
            rstfmt
            shfmt

            # Linters.
            golangci-lint
            shellcheck
          ];
        in {
          default = devenv.lib.mkShell {
            inherit inputs pkgs;
            modules = [{
              packages = codeGenPackages ++ languagePkgs ++ devPackages;

              languages.go = {
                enable = true;
                package = pkgs.go;
              };

              dotenv.enable = true;

              enterShell = ''
                export PATH="$(git rev-parse --show-toplevel)/dev/bin/:$PATH"
              '';
            }];
          };
        });
    };
}

