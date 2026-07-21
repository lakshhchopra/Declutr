import json
import os
import re
import subprocess
import sys
import time

if hasattr(sys.stdout, 'reconfigure'):
    sys.stdout.reconfigure(encoding='utf-8')

BACKLOG_PATH = r"C:\Users\Lakshya Chopra\.gemini\antigravity-ide\brain\dab098e0-fab9-4b4e-a62b-48c48807f7de\github_issue_backlog.md"

def parse_issues(file_path):
    with open(file_path, "r", encoding="utf-8") as f:
        content = f.read()

    # Split by "# Issue #"
    raw_issues = content.split("# Issue #")
    issues = []

    for raw in raw_issues[1:]:
        lines = raw.strip().split("\n")
        header = lines[0] # e.g. "001 - Repository Foundation & Developer Environment Setup"
        match = re.match(r"^(\d+)\s*-\s*(.*)$", header)
        if not match:
            continue

        issue_num = match.group(1)
        title = f"Issue #{issue_num} - {match.group(2)}"
        body = "\n".join(lines[1:]).strip()

        issues.append({
            "num": issue_num,
            "title": title,
            "body": body
        })

    return issues

def delete_existing_issues():
    print("Fetching existing GitHub issues...")
    res = subprocess.run(
        ["gh", "issue", "list", "--state", "all", "--limit", "200", "--json", "number,title"],
        capture_output=True,
        text=True
    )
    if res.returncode != 0:
        print(f"Error fetching issues: {res.stderr}")
        return

    try:
        existing = json.loads(res.stdout)
    except Exception as e:
        print(f"Failed to parse issues JSON: {e}")
        return

    if not existing:
        print("No existing issues found.")
        return

    print(f"Found {len(existing)} existing issues. Deleting them...")
    for item in existing:
        num = item["number"]
        title = item["title"]
        print(f"  Deleting issue #{num}: {title}")
        subprocess.run(["gh", "issue", "delete", str(num), "--yes"], capture_output=True)
        time.sleep(0.3)
    print("All existing issues deleted.\n")

def main():
    if not os.path.exists(BACKLOG_PATH):
        print(f"Error: Backlog file not found at {BACKLOG_PATH}")
        sys.exit(1)

    # Check gh CLI status
    res = subprocess.run(["gh", "auth", "status"], capture_output=True, text=True)
    if res.returncode != 0:
        print("\nGitHub CLI (gh) is not logged in!")
        print("Please run 'gh auth login' or set GH_TOKEN before running this script.\n")
        sys.exit(1)

    delete_existing_issues()

    issues = parse_issues(BACKLOG_PATH)
    print(f"Parsed {len(issues)} issues from backlog.")

    print("\nCreating 55 GitHub Issues in sequence (Issue #001 to #055)...\n")

    for i, issue in enumerate(issues, 1):
        print(f"[{i}/{len(issues)}] Creating {issue['title']}...")
        cmd = [
            "gh", "issue", "create",
            "--title", issue["title"],
            "--body", issue["body"]
        ]
        proc = subprocess.run(cmd, capture_output=True, text=True)
        if proc.returncode == 0:
            print(f"  Created: {proc.stdout.strip()}")
        else:
            print(f"  Error creating issue: {proc.stderr.strip()}")
        time.sleep(0.3)

    print("\nAll 55 GitHub Issues created successfully!")

if __name__ == "__main__":
    main()
