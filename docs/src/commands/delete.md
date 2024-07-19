# delete

Permanently deletes a task.

```bash
todo 12 delete
```

Deleting is not typically necessary as the `done` command will remove the task
from the list of uncompleted tasks. Deleting is only required if you want to
fully remove the task from history and reports.

Due to it's destructive nature, you will be prompted to confirm that you want to
delete the task(s) you specified. You can skip the prompt by adding the `-y`
flag.

```bash
todo delete -y
```

In addition to filtering by the task id, you can filter by project, priority,
tag, or task title.


```bash
todo +repair delete
```

Refer to the [filters](../filters.md) page for more details about the available
filters and how to use them effectively.
