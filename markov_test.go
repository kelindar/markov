package markov

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuilder(t *testing.T) {

	builder := New(2)
	builder.Teach("aac", "abc", "bbc", "aac", "abb", "bbb")

	model := builder.Build()
	assert.NotNil(t, model)

	r := rand.New(rand.NewSource(42))
	output := model.Generate(r)
	assert.Equal(t, "abbb", output)
}

func TestStars(t *testing.T) {

	builder := New(4)
	builder.Teach(stars...)

	model := builder.Build()
	assert.NotNil(t, model)

	r := rand.New(rand.NewSource(42))
	unique := 0
	for i := 0; i < 100; i++ {
		output := model.Generate(r)
		assert.NotEmpty(t, output)
		if !contains(stars, output) {
			unique++
		}
	}

	assert.NotZero(t, unique)
}

var stars = []string{"acamar", "achernar", "achird", "acrab", "akrab", "elakrab", "graffias", "acrux", "acubens", "adhafera", "adhara", "ain", "aladfar", "alamak", "alathfar", "alaraph", "albaldah", "albali", "albireo", "alchiba", "alcor", "alcyone", "aldebaran", "alderamin", "aldhafera", "aldhanab", "aldhibah", "aldib", "alfawaris", "alfecca", "meridiana", "alfirk", "algedi", "algiedi", "algenib", "algieba", "algol", "algorab", "alhajoth", "alhena", "alioth", "alkaid", "alkurud", "alkalb", "alrai", "alkalurops", "alkaphrah", "alkes", "alkurah", "almach", "alminliar", "alasad", "alnair", "alnasl", "alnilam", "alnitak", "alniyat", "alphard", "alphecca", "alpheratz", "alrai", "alrakis", "alrami", "alrischa", "alruccbah", "alsafi", "alsciaukat", "alshain", "alshat", "altair", "altais", "altarf", "alterf", "althalimain", "aludra", "alula", "alwaid", "alya", "alzir", "ancha", "angetenar", "ankaa", "antares", "arcturus", "arich", "arided", "arkab", "armus", "arneb", "arrakis", "ascella", "asellus", "ashlesha", "askella", "aspidiske", "asterion", "asterope", "atik", "atlas", "atria", "auva", "avior", "azaleh", "azelfafage", "azha", "azimech", "azmidiske", "baham", "biham", "baten", "kaitos", "becrux", "beid", "bellatrix", "benetnasch", "betelgeuse", "botein", "brachium", "canopus", "capella", "caph", "chaph", "caphir", "caput", "medusae", "castor", "castula", "cebalrai", "celbalrai", "ceginus", "celaeno", "chara", "cheleb", "chertan", "chort", "cor", "caroli", "hydrae", "leonis", "scorpii", "serpentis", "coxa", "cujam", "caiam", "cursa", "cynosura", "dabih", "decrux", "deneb", "denebola", "dheneb", "diadem", "diphda", "dnoces", "dschubba", "dubhe", "duhr", "edasich", "electra", "elmuthalleth", "elnath", "eltanin", "enif", "errai", "etamin", "fomalhaut", "fumalsamakah", "furud", "gacrux", "garnet", "gatria", "gemma", "gianfar", "giedi", "gienahgurab", "girtab", "gomeisa", "gorgonea", "graffias", "grafias", "grassias", "grumium", "hadar", "hadir", "haedus", "haldus", "hamal", "hassaleh", "hydrus", "heka", "heze", "hoedus", "homam", "hyadum", "hydrobius", "izar", "jabbah", "jih", "kabdhilinan", "kaffaljidhma", "kajam", "kastra", "kaus", "media", "keid", "kitalpha", "kleeia", "kochab", "kornephoros", "kraz", "ksora", "kullat", "kuma", "lanx", "librae", "superba", "lesath", "lucida", "maasym", "mahasim", "maia", "marfark", "marfik", "markab", "matar", "mebsuta", "megrez", "meissa", "mekbuda", "menchib", "menkab", "menkalinan", "menkar", "menkent", "menkib", "merak", "merga", "merope", "mesarthim", "miaplacidus", "mimosa", "minchir", "minelava", "minkar", "mintaka", "mira", "mirach", "miram", "mirfak", "mirzam", "misam", "mizar", "mothallah", "muliphein", "muphrid", "murzim", "muscida", "muscida", "muscida", "nair", "naos", "nash", "nashira", "navi", "nekkar", "nembus", "neshmet", "nihal", "nunki", "nusakan", "okul", "peacock", "phact", "phad", "pherkad", "pleione", "polaris", "pollux", "porrima", "praecipua", "procyon", "propus", "proxi", "pulcherrim", "rana", "ras", "rasalas", "rastaban", "thaoum", "regor", "regulus", "rigel", "rigil", "rijl", "rotanev", "ruchba", "ruchbah", "rukbat", "sabik", "sadachbia", "sadalbari", "sadalmelik", "sadalsuud", "sadatoni", "sadira", "sadr", "sadlamulk", "saiph", "saiph", "salm", "sargas", "sarin", "sceptrum", "scheat", "scheddi", "schedar", "segin", "seginus", "sham", "shaula", "sheliak", "sheratan", "shurnarkabti", "shashutu", "sinistra", "sirius", "situla", "skat", "spica", "sterope", "sterope", "sualocin", "subra", "suhail", "suhel", "sulafat", "sol", "syrma", "tabit", "talitha", "tania", "tarazet", "tarazed", "taygeta", "tegmen", "tegmine", "terebellum", "tejat", "thabit", "theemin", "thuban", "tien", "toliman", "torcularis", "tseen", "turais", "tyl", "unukalhai", "unuk", "vega", "vindemiatrix", "wasat", "wei", "wezen", "wezn", "yed", "yildun", "zaniah", "zaurak", "zavijava", "zawiat", "zedaron", "zelphah", "zibal", "zosma", "zuben", "zubenelgenubi", "zubeneschamali"}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
