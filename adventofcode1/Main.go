package test

import (
	"fmt"
	"learninggo/Utils"
	"learninggo/adventofcode1"
)

func main() {
	fmt.Println("Fetching Product of 2 numbers.....")
	var inputReport []int
	//inputReport = adventofcode1.IntializeInput()
	inputReport = Utils.IntializeInputInt("InputProblem1.txt")
	adventofcode1.GetProductTwoNumbers(inputReport)
	fmt.Println("Fetching Product of 3 numbers.....")
	adventofcode1.GetProductThreeNumbers(inputReport)
}

/*func total() bool {
	var inputEntries [6]int
	var sizeOfArray int
	var check bool

	fmt.Println("Enter Total number of entries:")
	fmt.Scanf("%d\n", &sizeOfArray)
	for i := 0; i < int(sizeOfArray); i++ {
		fmt.Println("Enter next value:")
		fmt.Scanf("%d\n", &inputEntries[i])
		fmt.Println(inputEntries[i])
	}

	for j := 0; j < int(sizeOfArray); j++ {
		for k := j + 1; k < int(sizeOfArray); k++ {
			check := inputEntries[j]+inputEntries[k] == 2020
			//fmt.Println(inputEntries[j])
			//fmt.Println(inputEntries[k])
			//fmt.Printf("Currently computing for %d and %d\n", &inputEntries[j], &inputEntries[k])
			//if inputEntries[j]+inputEntries[k] == 2020 {
			if check {
				fmt.Println(" The values matching the criteria are:")
				fmt.Println(inputEntries[j])
				fmt.Println(inputEntries[k])
			}
		}
	}
	return check
}*/

/*func getProduct() int {
	var inputEntries [6]int
	var sizeOfArray int
	var product int

	fmt.Println("Enter Total number of entries:")
	fmt.Scanf("%d\n", &sizeOfArray)
	for i := 0; i < int(sizeOfArray); i++ {
		fmt.Println("Enter next value:")
		fmt.Scanf("%d\n", &inputEntries[i])
	}

	for j := 0; j < int(sizeOfArray); j++ {
		for k := j + 1; k < int(sizeOfArray); k++ {
			check := inputEntries[j]+inputEntries[k] == 2020
			//fmt.Println(inputEntries[j])
			//fmt.Println(inputEntries[k])
			//fmt.Printf("Currently computing for %d and %d\n", &inputEntries[j], &inputEntries[k])
			//if inputEntries[j]+inputEntries[k] == 2020 {
			if check {
				product := inputEntries[j] * inputEntries[k]
				fmt.Println(" The values matching the criteria are:")
				fmt.Println(inputEntries[j])
				fmt.Println(inputEntries[k])
				fmt.Println(" The product of the values is:")
				fmt.Println(product)
			}
		}
	}
	return product
}*/

/*func getProduct() int {
	inputEntries := [200]int{1765, 1742, 1756, 1688, 1973, 1684, 1711, 1728, 1603, 1674, 1850, 1836, 1719, 1937, 1970, 1770, 1954, 1848, 1885, 1851, 1474, 1801, 1769, 1904, 1906, 1739, 1717, 1830, 1985, 1930, 1791, 1977, 1713, 1787, 1773, 1672, 1750, 1931, 1807, 1762, 1835, 1736, 1979, 1923, 1782, 1797, 1822, 1903, 1729, 343, 1678, 1753, 1873, 1358, 1987, 1821, 1761, 1988, 1886, 1669, 857, 1894, 1404, 1909, 1789, 1752, 1882, 1969, 1892, 1701, 1956, 1839, 483, 1897, 1730, 1829, 1879, 1810, 1755, 1799, 1932, 1715, 1809, 418, 1896, 1691, 1749, 1922, 1631, 1780, 1734, 1859, 1695, 1548, 1948, 1997, 1921, 1994, 1369, 1364, 1764, 1697, 1833, 1239, 616, 1786, 1890, 677, 1867, 1705, 1993, 1925, 1774, 1732, 1686, 1847, 1911, 1841, 1962, 1907, 1919, 1725, 1687, 1236, 1864, 1855, 1928, 1941, 1709, 1683, 1676, 1889, 1982, 1595, 1735, 1857, 1731, 1816, 2003, 1724, 741, 1655, 1308, 1959, 1955, 1768, 1795, 1804, 1961, 1693, 1884, 1813, 1927, 1845, 1689, 1646, 1803, 2008, 1599, 1984, 1871, 1814, 1918, 1990, 1858, 1908, 1949, 1980, 1618, 1718, 1712, 1989, 1876, 1947, 1974, 1838, 1865, 1842, 1817, 680, 1986, 1812, 1895, 1991, 1644, 1877, 1880, 1792, 1800, 1899, 1806, 1699, 1685, 1793, 1647, 1429, 1751, 1722, 1887, 1968}
	var product int
	//inputEntries[]  := {1778,1845,1813,1889,1939,1635,1443,796,1799,938,1488,1922,1909,1258,1659,1959,1557,1085,1379,1174,1782,1482,1702,1180,1992,1815,1802,215,1649,782,1847,1673,1823,1836,1447,1603,1767,1891,1964,1881,1637,1229,1994,1901,1583,1918,1415,1666,1155,1446,1315,1345,1948,1427,1242,1088,807,1747,1514,1351,1791,1612,1550,1926,1455,85,1594,1965,1884,1677,1960,1631,1585,1472,1263,1566,1998,1698,1968,1927,1378,1346,1710,1921,1827,1869,1187,1985,1323,1225,1474,1179,1580,1098,1737,1483,1665,1445,1979,1754,1854,1897,1405,1912,1614,1390,1773,1493,1333,1758,1867,1586,1347,1723,1285,394,1743,1252,320,1547,1804,1899,1526,1739,1533,1938,1081,1465,1920,1265,1470,1792,1118,1842,1204,1760,1663,893,1853,1244,1256,1428,1334,1967,1249,1752,1124,1725,1949,1340,1205,1584,548,1947,2002,1993,1931,1236,1154,1572,1650,1678,1944,1868,1129,1911,1106,1900,1240,1955,1219,1893,1459,1556,1173,1924,1568,1950,1303,1886,1365,1402,1711,1706,1671,1866,1403,1816,1717,1674,1487,1840,1951,1255,1786,1111,1280,1625,1478,1453}

	for j := 0; j < int(200); j++ {
		for k := j + 1; k < int(200); k++ {
			for l := k + 1; l < int(200); l++ {
				check := inputEntries[j]+inputEntries[k]+inputEntries[l] == 2020
				//fmt.Println(inputEntries[j])
				//fmt.Println(inputEntries[k])
				//fmt.Printf("Currently computing for %d and %d\n", &inputEntries[j], &inputEntries[k])
				//if inputEntries[j]+inputEntries[k] == 2020 {
				if check {
					product := inputEntries[j] * inputEntries[k] * inputEntries[l]
					fmt.Println(" The values matching the criteria are:")
					fmt.Println(inputEntries[j])
					fmt.Println(inputEntries[k])
					fmt.Println(" The product of the values is:")
					fmt.Println(product)
				}
			}
		}
	}
	return product
}*/

/*func getProductRecursion() int {
	inputEntries := [200]int{1765, 1742, 1756, 1688, 1973, 1684, 1711, 1728, 1603, 1674, 1850, 1836, 1719, 1937, 1970, 1770, 1954, 1848, 1885, 1851, 1474, 1801, 1769, 1904, 1906, 1739, 1717, 1830, 1985, 1930, 1791, 1977, 1713, 1787, 1773, 1672, 1750, 1931, 1807, 1762, 1835, 1736, 1979, 1923, 1782, 1797, 1822, 1903, 1729, 343, 1678, 1753, 1873, 1358, 1987, 1821, 1761, 1988, 1886, 1669, 857, 1894, 1404, 1909, 1789, 1752, 1882, 1969, 1892, 1701, 1956, 1839, 483, 1897, 1730, 1829, 1879, 1810, 1755, 1799, 1932, 1715, 1809, 418, 1896, 1691, 1749, 1922, 1631, 1780, 1734, 1859, 1695, 1548, 1948, 1997, 1921, 1994, 1369, 1364, 1764, 1697, 1833, 1239, 616, 1786, 1890, 677, 1867, 1705, 1993, 1925, 1774, 1732, 1686, 1847, 1911, 1841, 1962, 1907, 1919, 1725, 1687, 1236, 1864, 1855, 1928, 1941, 1709, 1683, 1676, 1889, 1982, 1595, 1735, 1857, 1731, 1816, 2003, 1724, 741, 1655, 1308, 1959, 1955, 1768, 1795, 1804, 1961, 1693, 1884, 1813, 1927, 1845, 1689, 1646, 1803, 2008, 1599, 1984, 1871, 1814, 1918, 1990, 1858, 1908, 1949, 1980, 1618, 1718, 1712, 1989, 1876, 1947, 1974, 1838, 1865, 1842, 1817, 680, 1986, 1812, 1895, 1991, 1644, 1877, 1880, 1792, 1800, 1899, 1806, 1699, 1685, 1793, 1647, 1429, 1751, 1722, 1887, 1968}
	//var product int
	n := 3
	sum := 0
	//inputEntries[]  := {1778,1845,1813,1889,1939,1635,1443,796,1799,938,1488,1922,1909,1258,1659,1959,1557,1085,1379,1174,1782,1482,1702,1180,1992,1815,1802,215,1649,782,1847,1673,1823,1836,1447,1603,1767,1891,1964,1881,1637,1229,1994,1901,1583,1918,1415,1666,1155,1446,1315,1345,1948,1427,1242,1088,807,1747,1514,1351,1791,1612,1550,1926,1455,85,1594,1965,1884,1677,1960,1631,1585,1472,1263,1566,1998,1698,1968,1927,1378,1346,1710,1921,1827,1869,1187,1985,1323,1225,1474,1179,1580,1098,1737,1483,1665,1445,1979,1754,1854,1897,1405,1912,1614,1390,1773,1493,1333,1758,1867,1586,1347,1723,1285,394,1743,1252,320,1547,1804,1899,1526,1739,1533,1938,1081,1465,1920,1265,1470,1792,1118,1842,1204,1760,1663,893,1853,1244,1256,1428,1334,1967,1249,1752,1124,1725,1949,1340,1205,1584,548,1947,2002,1993,1931,1236,1154,1572,1650,1678,1944,1868,1129,1911,1106,1900,1240,1955,1219,1893,1459,1556,1173,1924,1568,1950,1303,1886,1365,1402,1711,1706,1671,1866,1403,1816,1717,1674,1487,1840,1951,1255,1786,1111,1280,1625,1478,1453}

	for i := 0; i < int(197); i++ {
		sum = newMethod(i+1, n, sum, inputEntries)
		if inputEntries[i]+sum == 2020 {
			fmt.Println("Yes!!!!")
			break
		}
	}
	return 1
}

func newMethod(i int, n int, sum int, inputEntries [200]int) int {
	if n == i {
		return sum + inputEntries[i]
	}
	return newMethod(i+1, n-i+1, sum+inputEntries[i], inputEntries)
}*/
