package go_where39

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

const shuffleValue = 1337

// Room77, Berlin
var room77Coords = LatLng{Lat: 52.49303704, Lng: 13.41792593}
var emptyWordList = []string{}
var tooLongWordList = []string{"slush", "extend", "until", "layer", "arch", "small"}
var room77Words = []string{"slush", "extend", "until", "layer", "arch"}
var room77Shuffled1337 = []string{"small", "extra", "unusual", "lazy", "accident"}

func TestFromWordsArgs(t *testing.T) {
	// test an empty word list
	_, err := FromWords(emptyWordList, 1)
	if !strings.HasPrefix(err.Error(), "word list must contain between") {
		t.Log("should return error when passing an empty array")
		t.Fail()
	}

	// test a to long word list
	_, err = FromWords(tooLongWordList, 1)
	if !strings.HasPrefix(err.Error(), "word list must contain between") {
		t.Log("should return error when passing an empty array")
		t.Fail()
	}
}

func TestFromWordsConversion(t *testing.T) {
	coords, err := FromWords(room77Words, 1)
	if err != nil {
		t.Log("Should not return an error: ", err)
		t.Fail()
	}

	if fmt.Sprintf("%.8f", coords.Lat) != fmt.Sprintf("%.8f", room77Coords.Lat) || fmt.Sprintf("%.8f", coords.Lng) != fmt.Sprintf("%.8f", room77Coords.Lng) {
		t.Log("Conversion results should be ", room77Coords)
		t.Log("Conversion results should are ", coords)
		t.Fail()
	}
}

func TestFromWordsConversionShuffled(t *testing.T) {
	coords, err := FromWords(room77Shuffled1337, shuffleValue)
	if err != nil {
		t.Log("Should not return an error: ", err)
		t.Fail()
	}

	if fmt.Sprintf("%.8f", coords.Lat) != fmt.Sprintf("%.8f", room77Coords.Lat) || fmt.Sprintf("%.8f", coords.Lng) != fmt.Sprintf("%.8f", room77Coords.Lng) {
		t.Log("Conversion results should be ", room77Shuffled1337)
		t.Log("Conversion results should are ", coords)
		t.Fail()
	}
}

func TestToWordsConversion(t *testing.T) {
	words, err := ToWords(room77Coords, 1)

	if err != nil {
		t.Log("Error running ToWords() ", err)
		t.Fail()

	}

	if !reflect.DeepEqual(words, room77Words) {
		t.Log("Conversion results should be ", room77Words)
		t.Log("Conversion results should are ", words)
		t.Fail()
	}
}

func TestToWordsConversionShuffled(t *testing.T) {
	words, err := ToWords(room77Coords, shuffleValue)

	if err != nil {
		t.Log("Error running ToWords() ", err)
		t.Fail()

	}

	if !reflect.DeepEqual(words, room77Shuffled1337) {
		t.Log("Conversion results should be ", room77Shuffled1337)
		t.Log("Conversion results should are ", words)
		t.Fail()
	}
}

var fullWordListShuffled1337 = []string{"correct", "gentle", "off", "size", "card", "circle", "abandon", "damp", "dirt", "early", "erosion", "farm", "follow", "ability", "hamster", "hungry", "issue", "later",
	"luggage", "mercy", "name", "able", "patch", "point", "pupil", "rely", "room", "sell", "about", "sphere", "subject", "term", "tourist", "uniform", "village", "window",
	"above", "castle", "caught", "century", "chaos", "chef", "choice", "absent", "claw", "climb", "cloud", "coconut", "column", "concert", "cook", "absorb", "cover", "crash",
	"cricket", "crowd", "crystal", "curtain", "abstract", "day", "decorate", "delay", "depend", "design", "devote", "diet", "absurd", "display", "dog", "door", "draw", "drive",
	"during", "abuse", "ecology", "either", "elevator", "employ", "endorse", "enlist", "entry", "access", "eternal", "excess", "exhaust", "expire", "eyebrow", "false", "accident",
	"feature", "fence", "figure", "finger", "fitness", "flee", "flush", "account", "fortune", "frame", "frost", "fury", "gap", "gate", "accuse", "giraffe", "glide", "goat", "gossip",
	"grass", "grocery", "guitar", "achieve", "have", "hedgehog", "high", "hockey", "hope", "hour", "acid", "ice", "illness", "improve", "indicate", "initial", "inquiry", "into",
	"acoustic", "jeans", "joy", "kangaroo", "kidney", "kiwi", "labor", "acquire", "lawsuit", "left", "length", "library", "limit", "load", "long", "across", "mad", "mammal",
	"maple", "mask", "maximum", "medal", "act", "method", "minor", "mix", "monitor", "morning", "movie", "mushroom", "action", "neck", "net", "night", "notable", "number",
	"observe", "actor", "olympic", "open", "order", "other", "own", "page", "panther", "actress", "payment", "pencil", "phone", "pig", "pistol", "please", "actual", "portion",
	"powder", "pretty", "prison", "project", "provide", "adapt", "put", "quit", "radar", "random", "raw", "receive", "refuse", "add", "rent", "rescue", "return", "ribbon",
	"ring", "road", "addict", "rubber", "saddle", "salt", "sausage", "scene", "screen", "second", "address", "session", "shed", "shiver", "shrimp", "siege", "simple", "adjust",
	"slab", "slim", "smile", "soap", "soldier", "sort", "spare", "admit", "spoil", "spy", "stage", "steel", "stock", "strike", "adult", "suggest", "supply", "suspect", "sweet",
	"syrup", "talk", "teach", "advance", "there", "throw", "time", "tobacco", "tomato", "topic", "advice", "tragic", "treat", "trip", "trust", "turkey", "twist", "uncle", "aerobic",
	"unveil", "urban", "utility", "van", "vendor", "vessel", "affair", "visual", "vote", "walnut", "water", "weather", "whale", "whisper", "afford", "wise", "wool", "wrestle",
	"you", "car", "carbon", "afraid", "cargo", "carpet", "carry", "cart", "case", "cash", "casino", "again", "casual", "cat", "catalog", "catch", "category", "cattle", "age",
	"cause", "caution", "cave", "ceiling", "celery", "cement", "census", "agent", "cereal", "certain", "chair", "chalk", "champion", "change", "agree", "chapter", "charge",
	"chase", "chat", "cheap", "check", "cheese", "ahead", "cherry", "chest", "chicken", "chief", "child", "chimney", "aim", "choose", "chronic", "chuckle", "chunk", "churn",
	"cigar", "cinnamon", "air", "citizen", "city", "civil", "claim", "clap", "clarify", "airport", "clay", "clean", "clerk", "clever", "click", "client", "cliff", "aisle",
	"clinic", "clip", "clock", "clog", "close", "cloth", "alarm", "clown", "club", "clump", "cluster", "clutch", "coach", "coast", "album", "code", "coffee", "coil", "coin",
	"collect", "color", "alcohol", "combine", "come", "comfort", "comic", "common", "company", "alert", "conduct", "confirm", "congress", "connect", "consider", "control",
	"convince", "alien", "cool", "copper", "copy", "coral", "core", "corn", "all", "cost", "cotton", "couch", "country", "couple", "course", "cousin", "alley", "coyote",
	"crack", "cradle", "craft", "cram", "crane", "allow", "crater", "crawl", "crazy", "cream", "credit", "creek", "crew", "almost", "crime", "crisp", "critic", "crop",
	"cross", "crouch", "alone", "crucial", "cruel", "cruise", "crumble", "crunch", "crush", "cry", "alpha", "cube", "culture", "cup", "cupboard", "curious", "current",
	"already", "curve", "cushion", "custom", "cute", "cycle", "dad", "damage", "also", "dance", "danger", "daring", "dash", "daughter", "dawn", "alter", "deal", "debate",
	"debris", "decade", "december", "decide", "decline", "always", "decrease", "deer", "defense", "define", "defy", "degree", "amateur", "deliver", "demand", "demise",
	"denial", "dentist", "deny", "depart", "amazing", "deposit", "depth", "deputy", "derive", "describe", "desert", "among", "desk", "despair", "destroy", "detail", "detect",
	"develop", "device", "amount", "diagram", "dial", "diamond", "diary", "dice", "diesel", "amused", "differ", "digital", "dignity", "dilemma", "dinner", "dinosaur",
	"direct", "analyst", "disagree", "discover", "disease", "dish", "dismiss", "disorder", "anchor", "distance", "divert", "divide", "divorce", "dizzy", "doctor", "document",
	"ancient", "doll", "dolphin", "domain", "donate", "donkey", "donor", "anger", "dose", "double", "dove", "draft", "dragon", "drama", "drastic", "angle", "dream", "dress",
	"drift", "drill", "drink", "drip", "angry", "drop", "drum", "dry", "duck", "dumb", "dune", "animal", "dust", "dutch", "duty", "dwarf", "dynamic", "eager", "eagle", "ankle",
	"earn", "earth", "easily", "east", "easy", "echo", "announce", "economy", "edge", "edit", "educate", "effort", "egg", "eight", "annual", "elbow", "elder", "electric",
	"elegant", "element", "elephant", "another", "elite", "else", "embark", "embody", "embrace", "emerge", "emotion", "answer", "empower", "empty", "enable", "enact", "end",
	"endless", "antenna", "enemy", "energy", "enforce", "engage", "engine", "enhance", "enjoy", "antique", "enough", "enrich", "enroll", "ensure", "enter", "entire", "anxiety",
	"envelope", "episode", "equal", "equip", "era", "erase", "erode", "any", "error", "erupt", "escape", "essay", "essence", "estate", "apart", "ethics", "evidence", "evil",
	"evoke", "evolve", "exact", "example", "apology", "exchange", "excite", "exclude", "excuse", "execute", "exercise", "appear", "exhibit", "exile", "exist", "exit", "exotic",
	"expand", "expect", "apple", "explain", "expose", "express", "extend", "extra", "eye", "approve", "fabric", "face", "faculty", "fade", "faint", "faith", "fall", "april", "fame",
	"family", "famous", "fan", "fancy", "fantasy", "arch", "fashion", "fat", "fatal", "father", "fatigue", "fault", "favorite", "arctic", "february", "federal", "fee", "feed", "feel",
	"female", "area", "festival", "fetch", "fever", "few", "fiber", "fiction", "field", "arena", "file", "film", "filter", "final", "find", "fine", "argue", "finish", "fire", "firm",
	"first", "fiscal", "fish", "fit", "arm", "fix", "flag", "flame", "flash", "flat", "flavor", "armed", "flight", "flip", "float", "flock", "floor", "flower", "fluid", "armor", "fly",
	"foam", "focus", "fog", "foil", "fold", "army", "food", "foot", "force", "forest", "forget", "fork", "around", "forum", "forward", "fossil", "foster", "found", "fox", "fragile",
	"arrange", "frequent", "fresh", "friend", "fringe", "frog", "front", "arrest", "frown", "frozen", "fruit", "fuel", "fun", "funny", "furnace", "arrive", "future", "gadget",
	"gain", "galaxy", "gallery", "game", "arrow", "garage", "garbage", "garden", "garlic", "garment", "gas", "gasp", "art", "gather", "gauge", "gaze", "general", "genius", "genre",
	"artefact", "genuine", "gesture", "ghost", "giant", "gift", "giggle", "ginger", "artist", "girl", "give", "glad", "glance", "glare", "glass", "artwork", "glimpse", "globe",
	"gloom", "glory", "glove", "glow", "glue", "ask", "goddess", "gold", "good", "goose", "gorilla", "gospel", "aspect", "govern", "gown", "grab", "grace", "grain", "grant", "grape",
	"assault", "gravity", "great", "green", "grid", "grief", "grit", "asset", "group", "grow", "grunt", "guard", "guess", "guide", "guilt", "assist", "gun", "gym", "habit",
	"hair", "half", "hammer", "assume", "hand", "happy", "harbor", "hard", "harsh", "harvest", "hat", "asthma", "hawk", "hazard", "head", "health", "heart", "heavy",
	"athlete", "height", "hello", "helmet", "help", "hen", "hero", "hidden", "atom", "hill", "hint", "hip", "hire", "history", "hobby", "attack", "hold", "hole",
	"holiday", "hollow", "home", "honey", "hood", "attend", "horn", "horror", "horse", "hospital", "host", "hotel", "attitude", "hover", "hub", "huge", "human",
	"humble", "humor", "hundred", "attract", "hunt", "hurdle", "hurry", "hurt", "husband", "hybrid", "auction", "icon", "idea", "identify", "idle", "ignore",
	"ill", "illegal", "audit", "image", "imitate", "immense", "immune", "impact", "impose", "august", "impulse", "inch", "include", "income", "increase", "index",
	"aunt", "indoor", "industry", "infant", "inflict", "inform", "inhale", "inherit", "author", "inject", "injury", "inmate", "inner", "innocent", "input", "auto",
	"insane", "insect", "inside", "inspire", "install", "intact", "interest", "autumn", "invest", "invite", "involve", "iron", "island", "isolate", "average", "item",
	"ivory", "jacket", "jaguar", "jar", "jazz", "jealous", "avocado", "jelly", "jewel", "job", "join", "joke", "journey", "avoid", "judge", "juice", "jump", "jungle",
	"junior", "junk", "just", "awake", "keen", "keep", "ketchup", "key", "kick", "kid", "aware", "kind", "kingdom", "kiss", "kit", "kitchen", "kite", "kitten", "away",
	"knee", "knife", "knock", "know", "lab", "label", "awesome", "ladder", "lady", "lake", "lamp", "language", "laptop", "large", "awful", "latin", "laugh", "laundry",
	"lava", "law", "lawn", "awkward", "layer", "lazy", "leader", "leaf", "learn", "leave", "lecture", "axis", "leg", "legal", "legend", "leisure", "lemon", "lend", "baby",
	"lens", "leopard", "lesson", "letter", "level", "liar", "liberty", "bachelor", "license", "life", "lift", "light", "like", "limb", "bacon", "link", "lion", "liquid",
	"list", "little", "live", "lizard", "badge", "loan", "lobster", "local", "lock", "logic", "lonely", "bag", "loop", "lottery", "loud", "lounge", "love", "loyal", "lucky",
	"balance", "lumber", "lunar", "lunch", "luxury", "lyrics", "machine", "balcony", "magic", "magnet", "maid", "mail", "main", "major", "make", "ball", "man", "manage",
	"mandate", "mango", "mansion", "manual", "bamboo", "marble", "march", "margin", "marine", "market", "marriage", "banana", "mass", "master", "match", "material", "math",
	"matrix", "matter", "banner", "maze", "meadow", "mean", "measure", "meat", "mechanic", "bar", "media", "melody", "melt", "member", "memory", "mention", "menu",
	"barely", "merge", "merit", "merry", "mesh", "message", "metal", "bargain", "middle", "midnight", "milk", "million", "mimic", "mind", "minimum", "barrel", "minute",
	"miracle", "mirror", "misery", "miss", "mistake", "base", "mixed", "mixture", "mobile", "model", "modify", "mom", "moment", "basic", "monkey", "monster", "month",
	"moon", "moral", "more", "basket", "mosquito", "mother", "motion", "motor", "mountain", "mouse", "move", "battle", "much", "muffin", "mule", "multiply", "muscle",
	"museum", "beach", "music", "must", "mutual", "myself", "mystery", "myth", "naive", "bean", "napkin", "narrow", "nasty", "nation",
	"nature", "near", "beauty", "need", "negative", "neglect", "neither", "nephew", "nerve", "nest", "because", "network", "neutral", "never",
	"news", "next", "nice", "become", "noble", "noise", "nominee", "noodle", "normal", "north", "nose", "beef", "note", "nothing", "notice", "novel",
	"now", "nuclear", "before", "nurse", "nut", "oak", "obey", "object", "oblige", "obscure", "begin", "obtain", "obvious", "occur", "ocean", "october",
	"odor", "behave", "offer", "office", "often", "oil", "okay", "old", "olive", "behind", "omit", "once", "one", "onion", "online", "only", "believe", "opera",
	"opinion", "oppose", "option", "orange", "orbit", "orchard", "below", "ordinary", "organ", "orient", "original", "orphan", "ostrich", "belt", "outdoor",
	"outer", "output", "outside", "oval", "oven", "over", "bench", "owner", "oxygen", "oyster", "ozone", "pact", "paddle", "benefit", "pair", "palace", "palm",
	"panda", "panel", "panic", "best", "paper", "parade", "parent", "park", "parrot", "party", "pass", "betray", "path", "patient", "patrol", "pattern", "pause",
	"pave", "better", "peace", "peanut", "pear", "peasant", "pelican", "pen", "penalty", "between", "people", "pepper", "perfect", "permit", "person", "pet",
	"beyond", "photo", "phrase", "physical", "piano", "picnic", "picture", "piece", "bicycle", "pigeon", "pill", "pilot", "pink", "pioneer", "pipe", "bid", "pitch",
	"pizza", "place", "planet", "plastic", "plate", "play", "bike", "pledge", "pluck", "plug", "plunge", "poem", "poet", "bind", "polar", "pole", "police", "pond",
	"pony", "pool", "popular", "biology", "position", "possible", "post", "potato", "pottery", "poverty", "bird", "power", "practice", "praise", "predict", "prefer",
	"prepare", "present", "birth", "prevent", "price", "pride", "primary", "print", "priority", "bitter", "private", "prize", "problem", "process", "produce", "profit",
	"program", "black", "promote", "proof", "property", "prosper", "protect", "proud", "blade", "public", "pudding", "pull", "pulp", "pulse", "pumpkin", "punch", "blame",
	"puppy", "purchase", "purity", "purpose", "purse", "push", "blanket", "puzzle", "pyramid", "quality", "quantum", "quarter", "question", "quick", "blast", "quiz",
	"quote", "rabbit", "raccoon", "race", "rack", "bleak", "radio", "rail", "rain", "raise", "rally", "ramp", "ranch", "bless", "range", "rapid", "rare", "rate", "rather",
	"raven", "blind", "razor", "ready", "real", "reason", "rebel", "rebuild", "recall", "blood", "recipe", "record", "recycle", "reduce", "reflect", "reform", "blossom",
	"region", "regret", "regular", "reject", "relax", "release", "relief", "blouse", "remain", "remember", "remind", "remove", "render", "renew", "blue", "reopen", "repair",
	"repeat", "replace", "report", "require", "blur", "resemble", "resist", "resource", "response", "result", "retire", "retreat", "blush", "reunion", "reveal", "review",
	"reward", "rhythm", "rib", "board", "rice", "rich", "ride", "ridge", "rifle", "right", "rigid", "boat", "riot", "ripple", "risk", "ritual", "rival", "river", "body",
	"roast", "robot", "robust", "rocket", "romance", "roof", "rookie", "boil", "rose", "rotate", "rough", "round", "route", "royal", "bomb", "rude", "rug", "rule",
	"run", "runway", "rural", "sad", "bone", "sadness", "safe", "sail", "salad", "salmon", "salon", "bonus", "salute", "same", "sample", "sand", "satisfy", "satoshi",
	"sauce", "book", "save", "say", "scale", "scan", "scare", "scatter", "boost", "scheme", "school", "science", "scissors", "scorpion", "scout", "scrap", "border",
	"script", "scrub", "sea", "search", "season", "seat", "boring", "secret", "section", "security", "seed", "seek", "segment", "select", "borrow", "seminar", "senior",
	"sense", "sentence", "series", "service", "boss", "settle", "setup", "seven", "shadow", "shaft", "shallow", "share", "bottom", "shell", "sheriff", "shield", "shift",
	"shine", "ship", "bounce", "shock", "shoe", "shoot", "shop", "short", "shoulder", "shove", "box", "shrug", "shuffle", "shy", "sibling", "sick", "side", "boy", "sight",
	"sign", "silent", "silk", "silly", "silver", "similar", "bracket", "since", "sing", "siren", "sister", "situate", "six", "brain", "skate", "sketch", "ski", "skill",
	"skin", "skirt", "skull", "brand", "slam", "sleep", "slender", "slice", "slide", "slight", "brass", "slogan", "slot", "slow", "slush", "small", "smart", "brave",
	"smoke", "smooth", "snack", "snake", "snap", "sniff", "snow", "bread", "soccer", "social", "sock", "soda", "soft", "solar", "breeze", "solid", "solution", "solve",
	"someone", "song", "soon", "sorry", "brick", "soul", "sound", "soup", "source", "south", "space", "bridge", "spatial", "spawn", "speak", "special", "speed", "spell",
	"spend", "brief", "spice", "spider", "spike", "spin", "spirit", "split", "bright", "sponsor", "spoon", "sport", "spot", "spray", "spread", "spring", "bring", "square",
	"squeeze", "squirrel", "stable", "stadium", "staff", "brisk", "stairs", "stamp", "stand", "start", "state", "stay", "steak", "broccoli", "stem", "step", "stereo", "stick",
	"still", "sting", "broken", "stomach", "stone", "stool", "story", "stove", "strategy", "street", "bronze", "strong", "struggle", "student", "stuff", "stumble", "style",
	"broom", "submit", "subway", "success", "such", "sudden", "suffer", "sugar", "brother", "suit", "summer", "sun", "sunny", "sunset", "super", "brown", "supreme", "sure",
	"surface", "surge", "surprise", "surround", "survey", "brush", "sustain", "swallow", "swamp", "swap", "swarm", "swear", "bubble", "swift", "swim", "swing", "switch",
	"sword", "symbol", "symptom", "buddy", "system", "table", "tackle", "tag", "tail", "talent", "budget", "tank", "tape", "target", "task", "taste", "tattoo", "taxi", "buffalo",
	"team", "tell", "ten", "tenant", "tennis", "tent", "build", "test", "text", "thank", "that", "theme", "then", "theory", "bulb", "they", "thing", "this", "thought", "three",
	"thrive", "bulk", "thumb", "thunder", "ticket", "tide", "tiger", "tilt", "timber", "bullet", "tiny", "tip", "tired", "tissue", "title", "toast", "bundle", "today", "toddler",
	"toe", "together", "toilet", "token", "bunker", "tomorrow", "tone", "tongue", "tonight", "tool", "tooth", "top", "burden", "topple", "torch", "tornado", "tortoise", "toss",
	"total", "burger", "toward", "tower", "town", "toy", "track", "trade", "traffic", "burst", "train", "transfer", "trap", "trash", "travel", "tray", "bus", "tree", "trend", "trial",
	"tribe", "trick", "trigger", "trim", "business", "trophy", "trouble", "truck", "true", "truly", "trumpet", "busy", "truth", "try", "tube", "tuition", "tumble", "tuna", "tunnel",
	"butter", "turn", "turtle", "twelve", "twenty", "twice", "twin", "buyer", "two", "type", "typical", "ugly", "umbrella", "unable", "unaware", "buzz", "uncover", "under", "undo",
	"unfair", "unfold", "unhappy", "cabbage", "unique", "unit", "universe", "unknown", "unlock", "until", "unusual", "cabin", "update", "upgrade", "uphold", "upon", "upper",
	"upset", "cable", "urge", "usage", "use", "used", "useful", "useless", "usual", "cactus", "vacant", "vacuum", "vague", "valid", "valley", "valve", "cage", "vanish", "vapor",
	"various", "vast", "vault", "vehicle", "velvet", "cake", "venture", "venue", "verb", "verify", "version", "very", "call", "veteran", "viable", "vibrant", "vicious", "victory",
	"video", "view", "calm", "vintage", "violin", "virtual", "virus", "visa", "visit", "camera", "vital", "vivid", "vocal", "voice", "void", "volcano", "volume", "camp", "voyage",
	"wage", "wagon", "wait", "walk", "wall", "can", "want", "warfare", "warm", "warrior", "wash", "wasp", "waste", "canal", "wave", "way", "wealth", "weapon", "wear", "weasel",
	"cancel", "web", "wedding", "weekend", "weird", "welcome", "west", "wet", "candy", "what", "wheat", "wheel", "when", "where", "whip", "cannon", "wide", "width", "wife", "wild",
	"will", "win", "canoe", "wine", "wing", "wink", "winner", "winter", "wire", "wisdom", "canvas", "wish", "witness", "wolf", "woman", "wonder", "wood", "canyon", "word", "work",
	"world", "worry", "worth", "wrap", "wreck", "capable", "wrist", "write", "wrong", "yard", "year", "yellow", "capital", "young", "youth", "zebra", "zero", "zone", "zoo", "captain"}
