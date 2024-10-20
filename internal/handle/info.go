package handle

import (
	"github.com/yeungon/tuhuebot/pkg/helpers"
	tele "gopkg.in/telebot.v3"
)

var (
	deparment_room     = "Một số phòng chức năng:\n\nPhòng y tế: EI5, EI6\nTổ dữ liệu: Dãy L, tầng 1\n Phòng Đào tạo và CTSV: dãy F, tầng 1\n"
	location_coso_2    = `ơ sở 2: Đường Võ Văn Kiệt - Phường An Tây - Thành phố Huế,`
	thesis_requirement = "https://tieuhoc.org/static/images/2024_quydinh_lam_khoaluan.png"
	master_plan        = "https://tieuhoc.org/master/2024_2025.jpg"
	wifi_password      = "Trường: dhsph19572010\nKhoa: TU16051996"
	sotay_sinhvien_url = "https://tieuhoc.org/vanban/quydinh/SOTAYSINHVIEN_2021_tieuhoc.pdf"
	dieukien_lam_btl   = `Điều kiện để sinh viên được làm BTL: 

1. SV được đăng ký thực hiện BTL sau khi đã tích luỹ tối thiểu 15 tín chỉ (TC), có điểm TBC tích lũy đạt từ 2,4 trở lên. SV thực hiện BTL phải tham gia học tập
chuyên cần và thực hiện các yêu cầu học tập của GV.
2. Được GV phụ trách học phần đề nghị và Tổ trưởng chuyên môn quản lí học phần duyệt.
3. Trong mỗi học kì, một SV chỉ được phép ực hiện tối đa 02 BTL.

Điều kiện để giảng viên hướng dẫn BTL
1. Để được tham gia hướng dẫn BTL, GV đã giảng dạy đại học từ 1 năm trở lên. GV dạy học phần nào, thì hướng dẫn và chấm BTL của học phần đó.
2. Trong một năm học hướng dẫn không quá 12 BTL; không tham gia hướng dẫn BTL của người thân (vợ, chồng, con; anh, chi, em ruột).`

	dieukien_lam_khoaluan = `Điều kiện để sinh viên được đăng ký làm KL: 
						⠀⋆˚✿˖°⋆˚✿˖°⋆˚✿˖°
a. Để được nhận làm KL, sinh viên cần hội đủ các điều kiện sau đây:

1. Tổng số tín chỉ tích lũy đạt tối thiểu 75% khối lượng kiến thức tích lũy/khối lượng chương trình đào tạo của ngành học (xếp hạng SV năm thứ tư) và điểm trung bình chung tích lũy phải đạt từ 2,80 trở lên. (trong chương trình đào tạo có tính số lượng TC của các HP ngoại ngữ không chuyên)

2. Đã thực hiện ít nhất 1 TL có kết quả đạt từ 8,0 điểm trở lên.

3 . Đã tích luỹ một HP chuyên môn có từ 2 TC trở lên liên quan đến chuyên ngành mà SV đăng ký đề tài KL và đạt kết quả từ 8,5 điểm trở lên. 4. Số HP còn nợ hoặc học lại trong các học kỳ trước đó không quá 02 và không vượt quá tổng số 5 TC. 

b. Sinh viên làm đề tài NCKH độc lập và đã nghiệm thu được ưu tiên chọn giao thực hiện KL.

c. Mỗi Khoa xét duyệt số lượng SV được làm KL theo ngành học và không vượt quá 50% tổng số SV của khóa học thuộc ngành xét.

d. Những trường hợp đặc biệt sẽ trình Hiệu trưởng quyết định.`
	dieukien_lamtieuluan = `
Điều kiện để sinh viên được làm tiểu luận (TL): 

1. SV được đăng ký thực hiện TL sau khi đã tích luỹ tối thiểu 30 TC, có điểm TBC tích lũy đạt từ 2,5 trở lên và tối đa chỉ có 1 HP (có từ 2 TC trở lên) chưa đạt tích lũy.

2. SV thực hiện TL phải tham gia học tập chuyên cần và thực hiện các yêu cầu học tập của GV.

3. Được GV phụ trách học phần đề nghị, Tổ trưởng chuyên môn duyệt và báo cáo cho Trưởng khoa.

4. Trong mỗi học kỳ, một SV chỉ được phép thực hiện 01 TL.

Quy định chi tiết tại: https://tieuhoc.org/vanban/quydinh/quy_dinh_lam_khoaluan_tieuluan_20191101081351_2505_qd_dhsp.pdf
`
	dieukien_xet_totnghiep = `
1) Hoàn thành đủ số tín chỉ, đúng số môn theo chương trình của khóa mình theo học.

2) Có chứng chỉ Giáo dục thể chất.

3) Có chứng chỉ Giáo dục quốc phòng.

4) Có chứng chỉ ngoại ngữ B1 hoặc tương đương. Sinh viên Lào, Campuchia được miễn chứng chỉ ngoại ngữ.

5) Làm đơn xét tốt nghiệp.
`
)

func Info(b *tele.Bot) {
	kehoachnamhoc := tele.InlineButton{
		Unique: "btn_callback_kehoachnamhoc",
		Text:   "Kế hoạch năm học",
		Data:   "button1_clicked",
	}

	sodotruong := tele.InlineButton{
		Unique: "btn_callback_sodotruong",
		Text:   "Sơ đồ trường",
		Data:   "button1_clicked",
	}

	// Define the second inline button
	sotay_sinhvien := tele.InlineButton{
		Unique: "btn_callback2_sotaysinhvien",
		Text:   "Sổ tay sinh viên",
		Data:   "button2_clicked",
	}

	matkhau_wifi := tele.InlineButton{
		Unique: "btn_callback_matkhau_wifi",
		Text:   "Mật khẩu Wifi",
		Data:   "button1_clicked",
	}

	quydinh_lam_btl := tele.InlineButton{
		Unique: "btn_callback2_quydinh_lam_btl",
		Text:   "Điều kiện làm bài tập lớn",
		Data:   "button2_clicked",
	}
	quydinh_lamtieuluan := tele.InlineButton{
		Unique: "btn_callback2_quydinh_lam_tieuluan",
		Text:   "Điều kiện làm tiểu luận",
		Data:   "button2_clicked",
	}

	// Define the second inline button
	quydinh_lamkhoaluan := tele.InlineButton{
		Unique: "btn_callback2_quydinh_lamkhoaluan",
		Text:   "Điều kiện làm khóa lụân",
		Data:   "button2_clicked",
	}

	// Define the second inline button
	quydinh_xet_totnghiep := tele.InlineButton{
		Unique: "btn_callback2_quydinh_xettotnghiep",
		Text:   "Điều kiện xét tốt nghiệp",
		Data:   "button2_clicked",
	}

	// Create the reply markup and add both buttons in a single row
	inlineKeys := &tele.ReplyMarkup{}
	inlineKeys.InlineKeyboard = [][]tele.InlineButton{
		{sodotruong, sotay_sinhvien}, // Row 1: Button 1
		{kehoachnamhoc, matkhau_wifi},
		{quydinh_lam_btl},
		{quydinh_lamtieuluan},
		{quydinh_lamkhoaluan},
		{quydinh_xet_totnghiep},
	}

	b.Handle("/info", func(c tele.Context) error {
		// Create the reply markup and add the button
		return c.Send("Một số thông tin hữu tích:\n\n", inlineKeys)

	})

	b.Handle("info", func(c tele.Context) error {
		// Create the reply markup and add the button
		return c.Send("Một số thông tin hữu tích:\n\n", inlineKeys)

	})

	b.Handle(&helpers.Info, func(c tele.Context) error {
		return c.Send("Một số thông tin hữu tích:\n\n", inlineKeys)
	})

	b.Handle(&sodotruong, func(c tele.Context) error {
		link := "https://tieuhoc.org/map/sodo.jpg"
		photo := &tele.Photo{File: tele.FromURL(link)}
		c.Send(photo)
		c.Send(deparment_room)
		return nil
	})

	b.Handle(&sotay_sinhvien, func(c tele.Context) error {
		return c.Send(sotay_sinhvien_url)
	})

	b.Handle(&matkhau_wifi, func(c tele.Context) error {
		return c.Send(wifi_password)
	})

	b.Handle(&kehoachnamhoc, func(c tele.Context) error {
		photo := &tele.Photo{File: tele.FromURL(master_plan)}
		return c.Send(photo)
	})

	b.Handle(&quydinh_lam_btl, func(c tele.Context) error {
		return c.Send(dieukien_lam_btl)

	})

	b.Handle(&quydinh_lamtieuluan, func(c tele.Context) error {
		return c.Send(dieukien_lamtieuluan)

	})

	b.Handle(&quydinh_lamkhoaluan, func(c tele.Context) error {
		return c.Send(dieukien_lam_khoaluan)
	})

	b.Handle(&quydinh_xet_totnghiep, func(c tele.Context) error {
		return c.Send(dieukien_xet_totnghiep)
	})

}
