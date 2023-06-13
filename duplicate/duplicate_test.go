package duplicate

import (
	"fmt"
	"strings"
	"testing"
)

func TestD(t *testing.T) {
	mockData := []string{"vfde_s_0176", "vfde_s_0009", "vfde_s_0008", "vfde_s_0138", "vfde_s_00c5", "vfde_s_011f", "vfde_s_0033", "vfde_s_016d",
		"vfde_s_00d1", "vfde_s_00ce", "vfde_s_0003", "vfde_s_001e", "vfde_s_00f4", "vfde_s_0181", "vfde_s_004e", "vfde_s_007e", "vfde_s_0101", "vfde_s_00d9",
		"vfde_s_0084", "vfde_s_0029", "vfde_s_00cc", "vfde_s_0083", "vfde_s_001d", "vfde_s_00ea", "vfde_s_00c7", "vfde_s_004b", "vfde_s_0016", "vfde_s_00fa",
		"vfde_s_00f3", "vfde_s_00da", "vfde_s_0142", "vfde_s_001c", "vfde_s_011d", "vfde_s_00ca", "vfde_s_0104", "vfde_s_003a", "vfde_s_0102", "vfde_s_003d",
		"vfde_s_0151", "vfde_s_0122", "vfde_s_0109", "vfde_s_010f", "vfde_s_0027", "vfde_s_0144", "vfde_s_0032", "vfde_s_0080", "vfde_s_017b", "vfde_s_00d6",
		"vfde_s_0180", "vfde_s_0057", "vfde_s_00fd", "vfde_s_001f", "vfde_s_0110", "vfde_s_0178", "vfde_s_0025", "vfde_s_0034", "vfde_s_014f", "vfde_s_0036",
		"vfde_s_0035", "vfde_s_00fb", "vfde_s_0014", "vfde_s_004c", "vfde_s_00f8", "vfde_s_0082", "vfde_s_010d", "vfde_s_013c", "vfde_s_0005", "vfde_s_00dc",
		"vfde_s_017f", "vfde_s_005c", "vfde_s_010b", "vfde_s_0092", "vfde_s_00e6", "vfde_s_00eb", "vfde_s_0026", "vfde_s_0174", "vfde_s_016e", "vfde_s_0022",
		"vfde_s_012c", "vfde_s_004d", "vfde_s_0048", "vfde_s_00f5", "vfde_s_0129", "vfde_s_0023", "vfde_s_0020", "vfde_s_010e", "vfde_s_00f0", "vfde_s_0038",
		"vfde_s_0001", "vfde_s_0042", "vfde_s_014a", "vfde_s_007b", "vfde_s_010a", "vfde_s_0075", "vfde_s_000e", "vfde_s_0004", "vfde_s_0047", "vfde_s_000b",
		"vfde_s_002d", "vfde_s_002e", "vfde_s_0168", "vfde_s_0070", "vfde_s_0112", "vfde_s_014d", "vfde_s_0177", "vfde_s_0041", "vfde_s_0105", "vfde_s_00d3",
		"vfde_s_0030", "vfde_s_00f6", "vfde_s_002f", "vfde_s_0002", "vfde_s_002c", "vfde_s_0140", "vfde_s_0021", "vfde_s_003e", "vfde_s_012f", "vfde_s_016f",
		"vfde_s_0085", "vfde_s_0046", "vfde_s_013b", "vfde_s_0135", "vfde_s_012e", "vfde_s_017c", "vfde_s_0007", "vfde_s_00df", "vfde_s_0039", "vfde_s_0024",
		"vfde_s_0134", "vfde_s_0076", "vfde_s_0175", "vfde_s_003c", "vfde_s_0103", "vfde_s_0018", "vfde_s_0156", "vfde_s_0171", "vfde_s_0152", "vfde_s_014b",
		"vfde_s_0062", "vfde_s_00fc", "vfde_s_00ff", "vfde_s_00cd", "vfde_s_0179", "vfde_s_013f", "vfde_s_003b", "vfde_s_000f", "vfde_s_0114", "vfde_s_0136",
		"vfde_s_00ef", "vfde_s_0019", "vfde_s_017a", "vfde_s_00ee", "vfde_s_000c", "vfde_s_00cb", "vfde_s_00fe", "vfde_s_0013", "vfde_s_013e", "vfde_s_005a",
		"vfde_s_00bc", "vfde_s_0012", "vfde_s_00f9",
	}

	fmt.Println(HasDuplicate(mockData))
}

func HasDuplicate(slice []string) bool {
	seen := make(map[string]bool)

	for _, num := range slice {
		// If the element is already in the map, it's a duplicate
		if seen[num] {
			return true
		}
		// Mark the element as seen
		seen[num] = true
	}

	return false
}

func TestRemoveText(t *testing.T) {
	input := `"vfde_s_0176","vfde_s_0009","vfde_s_0008","vfde_s_0138","vfde_s_00c5","vfde_s_011f","vfde_s_0033","vfde_s_016d",
	"vfde_s_00d1","vfde_s_00ce","vfde_s_0003","vfde_s_001e","vfde_s_00f4","vfde_s_0181","vfde_s_004e","vfde_s_007e","vfde_s_0101","vfde_s_00d9",
	"vfde_s_0084","vfde_s_0029","vfde_s_00cc","vfde_s_0083","vfde_s_001d","vfde_s_00ea","vfde_s_00c7","vfde_s_004b","vfde_s_0016","vfde_s_00fa",
	"vfde_s_00f3","vfde_s_00da","vfde_s_0142","vfde_s_001c","vfde_s_011d","vfde_s_00ca","vfde_s_0104","vfde_s_003a","vfde_s_0102","vfde_s_003d",
	"vfde_s_0151","vfde_s_0122","vfde_s_0109","vfde_s_010f","vfde_s_0027","vfde_s_0144","vfde_s_0032","vfde_s_0080","vfde_s_017b","vfde_s_00d6",
	"vfde_s_0180","vfde_s_0057","vfde_s_00fd","vfde_s_001f","vfde_s_0110","vfde_s_0178","vfde_s_0025","vfde_s_0034","vfde_s_014f","vfde_s_0036",
	"vfde_s_0035","vfde_s_00fb","vfde_s_0014","vfde_s_004c","vfde_s_00f8","vfde_s_0082","vfde_s_010d","vfde_s_013c","vfde_s_0005","vfde_s_00dc",
	"vfde_s_017f","vfde_s_005c","vfde_s_010b","vfde_s_0092","vfde_s_00e6","vfde_s_00eb","vfde_s_0026","vfde_s_0174","vfde_s_016e","vfde_s_0022",
	"vfde_s_012c","vfde_s_004d","vfde_s_0048","vfde_s_00f5","vfde_s_0129","vfde_s_0023","vfde_s_0020","vfde_s_010e","vfde_s_00f0","vfde_s_0038",
	"vfde_s_0001","vfde_s_0042","vfde_s_014a","vfde_s_007b","vfde_s_010a","vfde_s_0075","vfde_s_000e","vfde_s_0004","vfde_s_0047","vfde_s_000b",
	"vfde_s_002d","vfde_s_002e","vfde_s_0168","vfde_s_0070","vfde_s_0112","vfde_s_014d","vfde_s_0177","vfde_s_0041","vfde_s_0105","vfde_s_00d3",
	"vfde_s_0030","vfde_s_00f6","vfde_s_002f","vfde_s_0002","vfde_s_002c","vfde_s_0140","vfde_s_0021","vfde_s_003e","vfde_s_012f","vfde_s_016f",
	"vfde_s_0085","vfde_s_0046","vfde_s_013b","vfde_s_0135","vfde_s_012e","vfde_s_017c","vfde_s_0007","vfde_s_00df","vfde_s_0039","vfde_s_0024",
	"vfde_s_0134","vfde_s_0076","vfde_s_0175","vfde_s_003c","vfde_s_0103","vfde_s_0018","vfde_s_0156","vfde_s_0171","vfde_s_0152","vfde_s_014b",
	"vfde_s_0062","vfde_s_00fc","vfde_s_00ff","vfde_s_00cd","vfde_s_0179","vfde_s_013f","vfde_s_003b","vfde_s_000f","vfde_s_0114","vfde_s_0136",
	"vfde_s_00ef","vfde_s_0019","vfde_s_017a","vfde_s_00ee","vfde_s_000c","vfde_s_00cb","vfde_s_00fe","vfde_s_0013","vfde_s_013e","vfde_s_005a",
	"vfde_s_00bc","vfde_s_0012","vfde_s_00f9"`
	result := strings.Replace(input, "\"", "", -1)
	substringsvalueStr := strings.Split(result, ",")
	outputValueStr := strings.Join(substringsvalueStr, ",")
	fmt.Println(outputValueStr)
}
