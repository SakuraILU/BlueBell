package rdb

type Script struct {
	Lua string
	Sha string
}

const (
	SETTOKEN = iota
	SETPOST
	SETVOTE
	GETVOTES
)

var scripts []Script = []Script{
	SETTOKEN: {
		Lua: `
			local key = KEYS[1]
			local token_str = ARGV[1]
			local cap = tonumber(ARGV[2])

			while redis.call("LLEN", key) >= cap do
				redis.call("LPOP", key)
			end

			redis.call("RPUSH", key, token_str)

			return 1
		`,
	},
	SETPOST: {
		Lua: `
			local post_score_key = KEYS[1]
			local post_time_key = KEYS[2]
			local pid = ARGV[1]
			local time = ARGV[2]

			redis.call("ZADD", post_score_key, 0, pid)
			redis.call("ZADD", post_time_key, time, pid)

			return 1
		`,
	},
	SETVOTE: {
		Lua: `
			local post_score_key = KEYS[1]
			local user_vote_post_key = KEYS[2]
			local pid = ARGV[1]
			local uid = ARGV[2]
			local orgchoice = tonumber(ARGV[3])
			local choice = tonumber(ARGV[4])

			local delta = choice - orgchoice
			redis.call("ZINCRBY", post_score_key, delta, pid)
			redis.call("ZADD", user_vote_post_key, choice, uid)

			return 1
		`,
	},
	GETVOTES: {
		Lua: `
			local nvotes = {}

			for _, key in ipairs(KEYS) do
				local vote, err = redis.call("ZCOUNT", key, ARGV[1], ARGV[1])
				
				if err then
					redis.log(redis.LOG_WARNING, "获取帖子的正面投票失败")
					return nil, err
				end

				table.insert(nvotes, vote)
			end

			return nvotes
		`,
	},
}
