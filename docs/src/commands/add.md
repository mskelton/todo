# add

Add a new task to the list.

```bash
todo add Buy milk
```

## Adding Tags

You can add any number of [tags](../tags.md) to your tasks to help organize
related tasks.

```bash
todo add Buy milk +shopping
todo add Buy groceries +shopping

# View all tasks with the shopping tag
todo +shopping list
```

## Set Priority

You can set the [priority](../priority.md) of a task to move it up or down in
your task list.

```bash
todo add Send thank you priority:H
```

To learn more how task order is determined, take a look at the [urgency](../urgency.md) section.

## Create a Recurring Task

[Recurring tasks](../recurrence.md) are a very important and powerful feature in todo. Modeled
closely after the recurrence options within Google Calendar, todo allows you to
express many different scenarios. For example, the following command will remind
you to water the plants every Tuesday.

```bash
todo add Water plants every:tues
```

Or if you need to walk the dog every Wednesday and Friday for the next 3 weeks:

```bash
todo add Walk the dog every:wed,fri until:3w
```

This barely scratches the surface of recurring tasks. There is so much more you
can do with recurring tasks which you can learn about on the
[recurrence](../recurrence.md) page.
