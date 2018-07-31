package entity

import "fmt"

type Area struct {
	Name string  `json:"name"`
	Slug string  `json:"slug"`
	Lat  float64 `json:"lat"`
	Lon  float64 `json:"lon"`
}

var (
	Dist1 = Area{
		Name: "Dist 1",
		Slug: "dist-1",
		Lat:  10.7766359,
		Lon:  106.6956761,
	}
	Dist2 = Area{
		Name: "Dist 2",
		Slug: "dist-2",
		Lat:  10.8009681,
		Lon:  106.7317152,
	}
	Dist3 = Area{
		Name: "Dist 3",
		Slug: "dist-3",
		Lat:  10.7788219,
		Lon:  106.6788454,
	}
	Dist4 = Area{
		Name: "Dist 4",
		Slug: "dist-4",
		Lat:  10.7596972,
		Lon:  106.7004623,
	}
	Dist5 = Area{
		Name: "Dist 5",
		Slug: "dist-5",
		Lat:  10.7549853,
		Lon:  106.6644184,
	}
	Dist6 = Area{
		Name: "Dist 6",
		Slug: "dist-6",
		Lat:  10.7456963,
		Lon:  106.6322585,
	}
	Dist7 = Area{
		Name: "Dist 7",
		Slug: "dist-7",
		Lat:  10.7330365,
		Lon:  106.7068752,
	}
	Dist8 = Area{
		Name: "Dist 8",
		Slug: "dist-8",
		Lat:  10.7395553,
		Lon:  106.6622515,
	}
	Dist9 = Area{
		Name: "Dist 9",
		Slug: "dist-9",
		Lat:  10.8182255,
		Lon:  106.7827821,
	}
	Dist10 = Area{
		Name: "Dist 10",
		Slug: "dist-10",
		Lat:  10.7694907,
		Lon:  106.6643915,
	}
	Dist11 = Area{
		Name: "Dist 11",
		Slug: "dist-11",
		Lat:  10.7627281,
		Lon:  106.6442081,
	}
	Dist12 = Area{
		Name: "Dist 12",
		Slug: "dist-12",
		Lat:  10.8720055,
		Lon:  106.6426596,
	}
	PhuNhuan = Area{
		Name: "Phu Nhuan",
		Slug: "phu-nhuan",
		Lat:  10.8001539,
		Lon:  106.6756577,
	}
	BinhThanh = Area{
		Name: "Binh Thanh",
		Slug: "binh-thanh",
		Lat:  10.802126,
		Lon:  106.7031109,
	}
	GoVap = Area{
		Name: "Go Vap",
		Slug: "go-vap",
		Lat:  10.8338809,
		Lon:  106.673991,
	}
	TanBinh = Area{
		Name: "Tan Binh",
		Slug: "tan-binh",
		Lat:  10.792969,
		Lon:  106.6466076,
	}
	TanPhu = Area{
		Name: "Tan Phu",
		Slug: "tan-phu",
		Lat:  10.7792864,
		Lon:  106.6235911,
	}
	ThuDuc = Area{
		Name: "Thu Duc",
		Slug: "thu-duc",
		Lat:  10.8469888,
		Lon:  106.7367636,
	}
	BinhTan = Area{
		Name: "Binh Tan",
		Slug: "binh-tan",
		Lat:  10.7912059,
		Lon:  106.5910111,
	}
)

// FindArea by name or slug
func FindArea(name string) (*Area, error) {
	switch name {
	case Dist1.Name, Dist1.Slug:
		return &Dist1, nil
	case Dist2.Name, Dist2.Slug:
		return &Dist2, nil
	case Dist3.Name, Dist3.Slug:
		return &Dist3, nil
	case Dist4.Name, Dist4.Slug:
		return &Dist4, nil
	case Dist5.Name, Dist5.Slug:
		return &Dist5, nil
	case Dist6.Name, Dist6.Slug:
		return &Dist6, nil
	case Dist7.Name, Dist7.Slug:
		return &Dist7, nil
	case Dist8.Name, Dist8.Slug:
		return &Dist8, nil
	case Dist9.Name, Dist9.Slug:
		return &Dist9, nil
	case Dist10.Name, Dist10.Slug:
		return &Dist10, nil
	case Dist11.Name, Dist11.Slug:
		return &Dist11, nil
	case Dist12.Name, Dist12.Slug:
		return &Dist12, nil
	case PhuNhuan.Name, PhuNhuan.Slug:
		return &PhuNhuan, nil
	case BinhThanh.Name, BinhThanh.Slug:
		return &BinhThanh, nil
	case GoVap.Name, GoVap.Slug:
		return &GoVap, nil
	case TanBinh.Name, TanBinh.Slug:
		return &TanBinh, nil
	case TanPhu.Name, TanPhu.Slug:
		return &TanPhu, nil
	case ThuDuc.Name, ThuDuc.Slug:
		return &ThuDuc, nil
	case BinhTan.Name, BinhTan.Slug:
		return &BinhTan, nil
	default:
		return nil, fmt.Errorf("unknown area %s", name)
	}
}

// ExistArea returns whether area name or slug is exists
func ExistArea(name string) bool {
	_, err := FindArea(name)
	return err == nil
}
