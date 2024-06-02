# @file setup git hooks for commitizen
# @brief Some shortcuts to setup git hooks
# @description
#   * git-cz for commit message
#   * pre-commit hooks

# @description new git-cz config in local repo
#
# @example
#   czmsg
#
# @stdout new changelog.config.js file
cfg_gcz() {
    cat <<EOF >changelog.config.js
module.exports = {
    disableEmoji: false,
    format: '{emoji}{type}{scope}: {subject}',
    list: ['feat', 'fix', 'docs', 'test', 'refactor', 'style', 'ci', 'perf', 'misc'],
    maxMessageLength: 64,
    minMessageLength: 3,
    questions: ['type', 'scope', 'subject', 'issues'],
    scopes: [],
    types: {
        misc: {
            description: 'Build process or auxiliary tool changes',
            emoji: '🍺',
            value: 'misc'
        },
        ci: {
            description: 'CI related changes',
            emoji: '🤖',
            value: 'ci'
        },
        docs: {
            description: 'Documentation only changes',
            emoji: '📝',
            value: 'docs'
        },
        feat: {
            description: 'A new feature',
            emoji: '✨',
            value: 'feat'
        },
        fix: {
            description: 'A bug fix',
            emoji: '🐛',
            value: 'fix'
        },
        perf: {
            description: 'A code change that improves performance',
            emoji: '⚡️',
            value: 'perf'
        },
        refactor: {
            description: 'A code change that neither fixes a bug or adds a feature',
            emoji: '💡',
            value: 'refactor'
        },
        release: {
            description: 'Create a release commit',
            emoji: '🔖',
            value: 'release'
        },
        style: {
            description: 'Markup, white-space, formatting, missing semi-colons...',
            emoji: '💄',
            value: 'style'
        },
        test: {
            description: 'Adding missing tests',
            emoji: '🧪',
            value: 'test'
        },
        messages: {
            type: 'Select the type of change that you\'re committing:',
            customScope: 'Select the scope this component affects:',
            subject: 'Write a short, imperative mood description of the change:\n',
            body: 'Provide a longer description of the change:\n ',
            breaking: 'List any breaking changes:\n',
            footer: 'Issues this commit closes, e.g #123:',
            confirmCommit: 'The packages that this commit has affected\n',
        },
    }
};
EOF
}

# @description new pre-commit config in local repo
#
# @example
#   new_precommit
#
# @stdout new .pre-commit-config.yaml file
cfg_precommit() {
    cat <<EOF >.pre-commit-config.yaml
repos:
- repo: https://github.com/gitleaks/gitleaks
  rev: v8.18.3
  hooks:
  - id: gitleaks
- repo: https://github.com/pre-commit/pre-commit-hooks
  rev: v4.6.0
  hooks:
  - id: check-added-large-files
  - id: check-json
  - id: check-yaml
  - id: end-of-file-fixer
  - id: trailing-whitespace
  # - id: check-merge-conflict
  # - id: check-xml
  # - id: check-toml
  # - id: file-contents-sorter
EOF

    pre-commit install
    pre-commit install-hooks
}

alias gcz="git cz"

alias gczfeat="git cz --type feat --subject"
alias gczdocs="git cz --type docs --subject"
alias gczfix="git cz --type fix --subject"
alias gcztest="git cz --type test --subject"
