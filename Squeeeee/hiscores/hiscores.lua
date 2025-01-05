local M = {}

local function get_file_path()
	return sys.get_save_file("squeeeee", "hiscores")
end

local data = {
	high_score = 0,
	games_played = 0,
	total_score = 0
}

function M.get_high_score()
	return data.high_score
end

function M.get_games_played()
	return data.games_played
end

function M.get_average_score()
	if data.games_played == 0 then
		return 0
	else
		return math.floor(data.total_score / data.games_played)
	end
end

local my_table = sys.load(get_file_path())
if next(my_table) then
	data.high_score = my_table.high_score
	data.games_played = my_table.games_played
	data.total_score = my_table.total_score
end

function M.record_game(score)
	data.high_score = math.max(score, data.high_score)
	data.games_played = data.games_played + 1
	data.total_score = data.total_score + score
	local my_table = {
		high_score = data.high_score,
		games_played = data.games_played,
		total_score = data.total_score
	}
	sys.save(get_file_path(), my_table)
end

return M