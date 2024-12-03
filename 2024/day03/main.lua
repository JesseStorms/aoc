-- i don't really use lua, but i wanna visualize this in pico-8 or retro gadgets
Token = {
    WHITESPACE=1,
    NUMBER=2,
    OPERATOR=3,
    LETTER=4,
    LEFT_PAREN=5,
    RIGHT_PAREN=6,
    COMMA=7,
    MISC=8,
    GUARD=9,
    DO=10,
    DONT=11
}

Token_type_labels = {
    [Token.WHITESPACE] = "WHITESPACE",
    [Token.NUMBER] = "NUMBER",
    [Token.OPERATOR] = "OPERATOR",
    [Token.LETTER] = "LETTER",
    [Token.LEFT_PAREN] = "LEFT_PAREN",
    [Token.RIGHT_PAREN] = "RIGHT_PAREN",
    [Token.COMMA] = "COMMA",
    [Token.MISC] = "MISC",
    [Token.GUARD] = "GUARD",
    [Token.DO] = "DO",
    [Token.DONT] = "DONT"
}

-- check if we are at the end of the text
function at_end(txt_info)
    return txt_info.pos > #txt_info.txt
end

-- see the next character in the text, but don't advance the position
function peek(txt_info)
    return txt_info.txt:sub(txt_info.pos, txt_info.pos)
end

-- advance the position in the text and return the character
function advance(txt_info)
    local ch = txt_info.txt:sub(txt_info.pos, txt_info.pos)
    txt_info.pos = txt_info.pos + 1
    return ch
end

-- check if the next characters in the text match the word
function matches(txt_info, word)
    -- bounds check
    if txt_info.pos + #word > #txt_info.txt then
        return false
    end
    -- check each character
    local i = txt_info.pos
    for j = 1, #word do
        local a = txt_info.txt:sub(i+j-1, i+j-1)
        local b = word:sub(j, j)
        if a ~= b then
            return false
        end
    end
    return true
end

-- scan a number from the text. Probably should make a scan word function too
function scan_number(txt_info, ch)
    local str = ch

    while peek(txt_info):match("%d") ~= nil do
        str = str .. advance(txt_info)
    end
    return str
end

-- figure out what a token is
function scan_token(txt_info)
    local tokens = {}
    while not at_end(txt_info) do
        local ch = advance(txt_info)
        
        if ch:match("%s") ~= nil then
            table.insert(tokens, {type=Token.WHITESPACE, value=ch})
        elseif ch:match("%d") ~= nil then
            table.insert(tokens, {type=Token.NUMBER, value=scan_number(txt_info, ch)})
        elseif ch == "(" then
            table.insert(tokens, {type=Token.LEFT_PAREN, value=ch})
        elseif ch == ")" then
            table.insert(tokens, {type=Token.RIGHT_PAREN, value=ch})
        elseif ch == "," then
            table.insert(tokens, {type=Token.COMMA, value=ch})
        elseif ch == "m" and matches(txt_info, "ul") then
            table.insert(tokens, {type=Token.OPERATOR, value="mul"})
            advance(txt_info)
            advance(txt_info)
        elseif ch == "d" and matches(txt_info, "o()") then
            table.insert(tokens, {type=Token.DO, value="do"})
            advance(txt_info)
            advance(txt_info)
            advance(txt_info)
        elseif ch == "d" and matches(txt_info, "on't()") then
            table.insert(tokens, {type=Token.DONT, value="don't"})
            advance(txt_info)
            advance(txt_info)
            advance(txt_info)
            advance(txt_info)
            advance(txt_info)
        elseif ch:match("%a") ~= nil then
            table.insert(tokens, {type=Token.LETTER, value=ch})
        else
            table.insert(tokens, {type=Token.MISC, value=ch})
        end
    end
    table.insert(tokens, {type=Token.GUARD, value=""}) -- reached the end
    return tokens
end
-- end of parsing stuff into tokens

-- get the next token in the instruction
function next_token(instruct)
    local t = instruct.tokens[instruct.pos]
    instruct.pos = instruct.pos + 1
    return t
end

-- peek at the next token in the instruction
function peek_token(instruct)
    return instruct.tokens[instruct.pos].type
end

-- doing the multiplication
function exec_mul(instruct, curr)
    if curr.type == Token.OPERATOR and peek_token(instruct) == Token.LEFT_PAREN then
        return exec_mul(instruct, next_token(instruct))
    elseif curr.type == Token.LEFT_PAREN and peek_token(instruct) == Token.NUMBER then
        return exec_mul(instruct, next_token(instruct))
    elseif curr.type == Token.NUMBER and peek_token(instruct) == Token.COMMA then
        instruct.a = curr.value
        return exec_mul(instruct, next_token(instruct))
    elseif curr.type == Token.COMMA and peek_token(instruct) == Token.NUMBER then
        return exec_mul(instruct, next_token(instruct))
    elseif curr.type == Token.NUMBER and peek_token(instruct) == Token.RIGHT_PAREN then
        instruct.b = curr.value
        return exec_mul(instruct, next_token(instruct))
    elseif curr.type == Token.RIGHT_PAREN then
        return instruct.a * instruct.b
    end
    return 0
end

--part 1 parsing
function exec1(instruct)
    local sum = 0
    while instruct.pos < #instruct.tokens do
        local t = next_token(instruct)
        if t.type == Token.GUARD then
            print("guard hit")
            break
        elseif t.type == Token.OPERATOR and t.value == "mul" then
            instruct.a = 0
            instruct.b = 0
            sum = sum + exec_mul(instruct, next_token(instruct))
        else 
            -- print("huh? " .. t.type)
        end
    end
    return sum
end

--part2 execution
function exec2(instruct)
    local sum = 0
    local working = true
    while instruct.pos < #instruct.tokens do
        local t = next_token(instruct)      
        if t.type == Token.GUARD then
            break
        elseif t.type == Token.OPERATOR and t.value == "mul" and working then
            instruct.a = 0
            instruct.b = 0
            sum = sum + exec_mul(instruct, next_token(instruct))
        elseif t.type == Token.DO then
            working = true
        elseif t.type == Token.DONT then
            working = false
        end
    end
    return sum
end


function read_file()
    local f = io.open("input.txt", "r")
    local txt = f:read("*all")
    f:close()
    local txt_info = {txt=txt, pos=1}
    local tokens = scan_token(txt_info)
    local instruct = {tokens=tokens, pos=1, a=0, b=0}
    return instruct
end

function part1()
    local instruct = read_file()
    return exec1(instruct)
end

function part2()
    local instruct = read_file()
    return exec2(instruct)
end

print(part1())
print(part2())