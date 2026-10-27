# Policy: Commits and Comments

## Commits

One commit per gesture. It holds the code and its tests.

Format (semantic release):

```
<type>(<ticket>): <gesture in Ubiquitous Language>
```

- `feat(CM-300): reject invoice when VAT number is invalid`
- `fix(CM-300): keep mandate active after a failed renewal`
- `refactor(CM-300): move VAT check into InvoiceIssuance`

Rules:
- Type: `feat`, `fix`, `refactor`, `test`, `docs` or `chore`.
- The subject says the behavior. Imperative, 72 characters or less, no final dot.
- No gate, agent or role name in the message.
- No ticket known: leave out the scope (`feat: ...`).
- A body is optional. Use it to say why.
- Commit locally with `--no-gpg-sign`. Never use `--no-verify`. If a hook fails, fix the cause.
- Never push. The human approves the push at the end.
- If `.ai/local/` sets another commit rule, follow it.

After the review, the human signs the commits in one step:

```sh
git rebase <base> --exec 'git commit --amend --no-edit -S'
```

Files in `documentation-agent/` are not committed by agents.

## Comments in code

- Short. Simple English. Explain why, not what.
- Self-contained: a reader with only the code understands them.
- Never mention a ticket, a gate, a role (PO, DEV-FS), or this workflow.
- No comment when the name already says it.
