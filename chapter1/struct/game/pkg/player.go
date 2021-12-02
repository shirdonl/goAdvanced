package pkg

type Player struct {
	CurrentVector Vector
	TargetVector  Vector
	Speed         float32
}

// 初始化玩家，设置速度
func NewPlayer(Speed float32) Player {
	return Player{Speed: Speed}
}

// 设置目标位置
func (p *Player) MoveTo(v Vector) {
	p.TargetVector = v
}

// 获取当前位置
func (p Player) Position() Vector {
	return p.CurrentVector
}

// 是否到达目标位置
func (p Player) IsArrived() bool {
	return p.CurrentVector.distanceTo(p.TargetVector) < p.Speed
}

// 更新玩家位置
func (p *Player) Update() {
	directionVector := p.TargetVector.sub(p.CurrentVector)
	normalizeVector := directionVector.normalize()
	pointChange := normalizeVector.multi(p.Speed)
	newVector := p.CurrentVector.add(pointChange)
	p.CurrentVector = newVector
}
