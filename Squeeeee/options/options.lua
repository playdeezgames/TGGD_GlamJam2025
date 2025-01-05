local M = {}

M.sfx_enabled = true

local function get_file_path()
	return sys.get_save_file("squeeeee", "options")
end

function M.load()
	local my_file_path = get_file_path()
	local my_table = sys.load(my_file_path)
	if next(my_table) then
		if my_table.sfx_enabled ~= nil then
			M.sfx_enabled = my_table.sfx_enabled
		end
	end
end

function M.save()
	local data = {}
	local my_file_path = get_file_path()
	data.sfx_enabled = M.sfx_enabled
	sys.save(my_file_path, data)
end

return M