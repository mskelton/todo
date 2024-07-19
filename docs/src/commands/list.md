# list

Show the task list.

```bash
todo list
```

Because `list` is the default command for todo, you can omit it from commands.

```bash
todo
```

_If you specify a single argument to todo which is the id of a task, the
[`show`](./show.md) command will be used instead of `list`. If you want to view
a single task in the list view you must include `list` as an argument._

## Filtering Tasks

You can specify [filters](../filters.md) to view a subset of tasks. For example,
use the following command to print all tasks with the `shopping` tag that are
high priority and contain the text "milk".

```bash
todo +shopping priority:H milk
```

Refer to the [filters](../filters.md) page for more details about the available
filters and how to use them effectively.
