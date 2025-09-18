package problems

type Task struct {
	u, t, p int
}
type TaskManager struct {
	idToTask map[int]Task
	q        []Task
}

func TaskLessThan(a, b Task) bool {
	if a.p == b.p {
		return b.t < a.t
	}
	return b.p < a.p
}

func THPush(h []Task, x Task) []Task {
	h = append(h, x)
	cur := len(h) - 1
	for cur > 0 {
		p := (cur - 1) >> 1
		if TaskLessThan(h[p], h[cur]) {
			break
		}
		h[p], h[cur] = h[cur], h[p]
		cur = p
	}
	return h
}

func THPop(h []Task) []Task {
	last := h[len(h)-1]
	h[0] = last
	h = h[:len(h)-1]
	cur, l := 0, 1
	for l < len(h) {
		r := l + 1
		c := l
		if r < len(h) && TaskLessThan(h[r], h[l]) {
			c = r
		}
		if TaskLessThan(h[cur], h[c]) {
			break
		}
		h[cur], h[c] = h[c], h[cur]
		cur = c
		l = cur*2 + 1
	}
	return h
}

func TaskManagerConstructor(tasks [][]int) TaskManager {
	idToTask := make(map[int]Task, len(tasks))
	q := make([]Task, 0)
	for _, task := range tasks {
		u, t, p := task[0], task[1], task[2]
		tt := Task{u: u, t: t, p: p}
		idToTask[t] = tt
		q = THPush(q, tt)
	}
	return TaskManager{idToTask: idToTask, q: q}
}

func (tm *TaskManager) Add(userId int, taskId int, priority int) {
	tt := Task{u: userId, t: taskId, p: priority}
	tm.idToTask[taskId] = tt
	tm.q = THPush(tm.q, tt)
}

func (tm *TaskManager) Edit(taskId int, newPriority int) {
	tt := tm.idToTask[taskId]
	if tt.p == newPriority {
		return
	}
	tt.p = newPriority
	tm.idToTask[taskId] = tt
	tm.q = THPush(tm.q, tt)
}

func (tm *TaskManager) Rmv(taskId int) {
	delete(tm.idToTask, taskId)
}

func (tm *TaskManager) ExecTop() int {
	for len(tm.q) > 0 {
		top := tm.q[0]
		tt, has := tm.idToTask[top.t]
		if has && tt.u == top.u && tt.p == top.p {
			break
		}
		tm.q = THPop(tm.q)
	}
	if len(tm.q) == 0 {
		return -1
	}
	top := tm.q[0]
	tm.q = THPop(tm.q)
	delete(tm.idToTask, top.t)
	return top.u
}

/**
 * Your TaskManager object will be instantiated and called as such:
 * obj := Constructor(tasks);
 * obj.Add(userId,taskId,priority);
 * obj.Edit(taskId,newPriority);
 * obj.Rmv(taskId);
 * param_4 := obj.ExecTop();
 */
