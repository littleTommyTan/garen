package dao

func (d *Dao) VerifyToken(token string) (string, error) {
	openId, err := d.Redis.Get(token).Result()
	return openId, err
}
