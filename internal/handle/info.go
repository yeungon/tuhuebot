package handle

import (
	"github.com/yeungon/tuhuebot/pkg/helpers"
	tele "gopkg.in/telebot.v3"
)

var (
	thesis_requirement    = "https://tieuhoc.org/static/images/2024_quydinh_lam_khoaluan.png"
	master_plan           = "https://tieuhoc.org/master/2024_2025.jpg"
	sotay_sinhvien        = "https://tieuhoc.org/vanban/quydinh/SOTAYSINHVIEN_2021_tieuhoc.pdf"
	dieukien_lam_khoaluan = `Điều kiện để sinh viên được đăng ký làm KL: 
						⠀⋆˚✿˖°⋆˚✿˖°⋆˚✿˖°
a. Để được nhận làm KL, sinh viên cần hội đủ các điều kiện sau đây:

1. Tổng số tín chỉ tích lũy đạt tối thiểu 75% khối lượng kiến thức tích lũy/khối lượng chương trình đào tạo của ngành học (xếp hạng SV năm thứ tư) và điểm trung bình chung tích lũy phải đạt từ 2,80 trở lên. (trong chương trình đào tạo có tính số lượng TC của các HP ngoại ngữ không chuyên)

2. Đã thực hiện ít nhất 1 TL có kết quả đạt từ 8,0 điểm trở lên.

3 . Đã tích luỹ một HP chuyên môn có từ 2 TC trở lên liên quan đến chuyên ngành mà SV đăng ký đề tài KL và đạt kết quả từ 8,5 điểm trở lên. 4. Số HP còn nợ hoặc học lại trong các học kỳ trước đó không quá 02 và không vượt quá tổng số 5 TC. 

b. Sinh viên làm đề tài NCKH độc lập và đã nghiệm thu được ưu tiên chọn giao thực hiện KL.

c. Mỗi Khoa xét duyệt số lượng SV được làm KL theo ngành học và không vượt quá 50% tổng số SV của khóa học thuộc ngành xét.

d. Những trường hợp đặc biệt sẽ trình Hiệu trưởng quyết định.`
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

	// Define the second inline button
	quydinh_lamkhoaluan := tele.InlineButton{
		Unique: "btn_callback2_quydinh_lamkhoaluan",
		Text:   "Điều kiện để được làm khóa lụân",
		Data:   "button2_clicked",
	}

	// Create the reply markup and add both buttons in a single row
	inlineKeys := &tele.ReplyMarkup{}
	inlineKeys.InlineKeyboard = [][]tele.InlineButton{
		{sodotruong},     // Row 1: Button 1
		{sotay_sinhvien}, // Row 2: Button 2
		{kehoachnamhoc},
		{quydinh_lamkhoaluan},
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
		return c.Send(photo)

	})

	b.Handle(&sotay_sinhvien, func(c tele.Context) error {
		return c.Send(sotay_sinhvien)
	})

	b.Handle(&kehoachnamhoc, func(c tele.Context) error {
		photo := &tele.Photo{File: tele.FromURL(master_plan)}
		return c.Send(photo)
	})

	b.Handle(&quydinh_lamkhoaluan, func(c tele.Context) error {
		//photo := &tele.Photo{File: tele.FromURL(thesis_requirement)}
		return c.Send(dieukien_lam_khoaluan)
		//return c.Send(photo)
	})

}
