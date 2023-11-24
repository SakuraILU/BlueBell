package ratelimit

var init_script = `
	local key_ntoken = KEYS[1]
	local key_lasttime = KEYS[2]

	local curtime = tonumber(ARGV[1])

	redis.call("SET", key_ntoken, 0)
	redis.call("SET", key_lasttime, curtime)

	return 1
`

var allow_script = `
	local key_ntoken = KEYS[1]
	local key_lasttime = KEYS[2]

	local need = tonumber(ARGV[1])
	local rate = tonumber(ARGV[2])
	local nburst = tonumber(ARGV[3])
	local curtime = tonumber(ARGV[4])

	local ntoken = tonumber(redis.call("GET", key_ntoken))

	local lasttime = tonumber(redis.call("GET", key_lasttime))
	local delta = (curtime - lasttime) * rate

	ntoken = ntoken + delta
	if ntoken > nburst then
		ntoken = nburst
	end

	if ntoken < need then  -- Changed from ntoken <= need
		return 0
	end

	ntoken = ntoken - need

	redis.call("SET", key_ntoken, ntoken)
	redis.call("SET", key_lasttime, curtime)

	return 1
`
