import os

EXCLUDE_DIRS = {'.git', '.next', 'node_modules', 'vendor', '.cache', 'dist', 'build', '__pycache__', '.idea', '.vscode'}

def build_tree(root_dir):
    lines = []
    lines.append("# Declutr Project — Full Directory Structure\n")
    lines.append("This document contains the complete folder structure of the **Declutr** repository across all 40 engineering milestones.\n")
    lines.append("```")
    
    root_dir = os.path.abspath(root_dir)
    root_name = os.path.basename(root_dir)
    lines.append(f"{root_name}/")

    def _walk(current_dir, prefix=""):
        try:
            entries = sorted(os.listdir(current_dir))
        except PermissionError:
            return
        
        filtered = [e for e in entries if e not in EXCLUDE_DIRS and not e.startswith('.sys')]
        for i, entry in enumerate(filtered):
            path = os.path.join(current_dir, entry)
            is_last = (i == len(filtered) - 1)
            connector = "└── " if is_last else "├── "
            
            if os.path.isdir(path):
                lines.append(f"{prefix}{connector}{entry}/")
                new_prefix = prefix + ("    " if is_last else "│   ")
                _walk(path, new_prefix)
            else:
                lines.append(f"{prefix}{connector}{entry}")

    _walk(root_dir)
    lines.append("```\n")
    return "\n".join(lines)

if __name__ == "__main__":
    repo_root = "f:\\Github\\Declutr"
    tree_md = build_tree(repo_root)
    out_path = os.path.join(repo_root, "docs", "declutr_project_structure.md")
    with open(out_path, "w", encoding="utf-8") as f:
        f.write(tree_md)
    print(f"Project structure written to {out_path}")
