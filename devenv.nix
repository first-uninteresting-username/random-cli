{ pkgs, ... }:

{
  packages = [ pkgs.git ];

  git-hooks.hooks = {
    check-added-large-files.enable = true;
    check-merge-conflicts.enable = true;
    detect-private-keys.enable = true;
    end-of-file-fixer = {
      enable = true;
      excludes = [ "devenv.lock" ];
    };
    trim-trailing-whitespace.enable = true;
    check-json.enable = true;
    check-yaml.enable = true;
    alejandra.enable = true;
  };
}
