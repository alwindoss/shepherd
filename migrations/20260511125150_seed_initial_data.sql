-- +goose Up
-- +goose StatementBegin

-- 1. Insert Roles
INSERT INTO roles (name, description) VALUES 
('Admin', 'Full system access, including user management'),
('Manager', 'Can view all donors and manage donations'),
('Staff', 'Data entry for donations and donor updates');

-- 2. Insert Initial System Users (Staff)
-- Note: password_hash should be a real hash in production
INSERT INTO users (username, first_name, last_name, email, password_hash) VALUES 
('shepherd', 'Jane', 'Doe', 'shepherd@shepherd.org', '$2a$10$InCxgtSllgvwFYC.qBsg1uhr.SyS5r1w.4xHJStxAIt/YfgVPZddy'),
('jdoe_admin', 'John', 'Doe', 'jdoe_admin@shepherd.org', '$2a$10$InCxgtSllgvwFYC.qBsg1uhr.SyS5r1w.4xHJStxAIt/YfgVPZddy'),
('samy_manager', 'Samy', 'Smith', 'samy@shepherd.org', '$2a$10$InCxgtSllgvwFYC.qBsg1uhr.SyS5r1w.4xHJStxAIt/YfgVPZddy');

-- 3. Assign Roles to Users
INSERT INTO user_roles (user_id, role_id) VALUES 
(1, 1), -- John is Admin
(2, 2); -- Samy is Manager

-- 4. Insert Supported Currencies
INSERT INTO currencies (code, symbol, name) VALUES 
('INR', '₹', 'Indian Rupee'),
('USD', '$', 'US Dollar'),
('EUR', '€', 'Euro');

-- 5. Insert Donors
INSERT INTO members (first_name, last_name, email, phone, aadhar_number, pan, address, requested_anonymity) VALUES 
('Aarav', 'Patel', 'aarav.patel.1@example.in', '919000000001', '100000000001', 'ABCDE0001A', '123 Lotus Lane, Mumbai', false),
('Ishani', 'Iyer', 'ishani.iyer.2@testmail.com', '919000000002', '100000000002', 'BCDEF0002B', '456 Temple St, Chennai', true),
('Vihaan', 'Singh', 'vihaan.s.3@domain.in', '919000000003', '100000000003', 'CDEFG0003C', '789 Park Ave, Delhi', false),
('Ananya', 'Reddy', 'ananya.r.4@provider.com', '919000000004', '100000000004', 'DEFGH0004D', '321 Hill View, Hyderabad', false),
('Arjun', 'Das', 'arjun.das.5@example.com', '919000000005', '100000000005', 'EFGHI0005E', '654 Lake Side, Kolkata', true),
('Saanvi', 'Nair', 'saanvi.n.6@mail.in', '919000000006', '100000000006', 'FGHIJ0006F', '987 Green Garden, Kochi', false),
('Sai', 'Kulkarni', 'sai.k.7@testmail.in', '919000000007', '100000000007', 'GHIJK0007G', '147 Heritage Colony, Pune', false),
('Aadhya', 'Gupta', 'aadhya.g.8@example.in', '919000000008', '100000000008', 'HIJKL0008H', '258 Central Plaza, Indore', true),
('Krishna', 'Bose', 'k.bose.9@domain.com', '919000000009', '100000000009', 'IJKLM0009I', '369 Victoria Rd, Kolkata', false),
('Myra', 'Chatterjee', 'myra.c.10@mail.com', '919000000010', '100000000010', 'JKLMN0010J', '741 Skyline Apts, Bengaluru', false),
('Aryan', 'Mehta', 'aryan.m.11@example.in', '919000000011', '100000000011', 'KLMNO0011K', '852 Ocean View, Ahmedabad', false),
('Avni', 'Joshi', 'avni.j.12@provider.in', '919000000012', '100000000012', 'LMNOP0012L', '963 Royal Enclave, Jaipur', true),
('Kabir', 'Pandey', 'kabir.p.13@testmail.com', '919000000013', '100000000013', 'MNOPQ0013M', '159 Ganga Vihar, Varanasi', false),
('Diya', 'Saxena', 'diya.s.14@example.com', '919000000014', '100000000014', 'NOPQR0014N', '357 Rose Wood, Lucknow', false),
('Rudra', 'Verma', 'rudra.v.15@mail.in', '919000000015', '100000000015', 'OPQRS0015O', '482 Silver Oak, Chandigarh', true),
('Ritika', 'Mishra', 'ritika.m.16@domain.in', '919000000016', '100000000016', 'PQRST0016P', '619 Diamond City, Surat', false),
('Reyansh', 'Yadav', 'reyansh.y.17@example.in', '919000000017', '100000000017', 'QRSTU0017Q', '274 Metro Heights, Gurgaon', false),
('Aarohi', 'Deshmukh', 'aarohi.d.18@testmail.in', '919000000018', '100000000018', 'RSTUV0018R', '593 Peace Villa, Pune', true),
('Vivaan', 'Malhotra', 'vivaan.m.19@mail.com', '919000000019', '100000000019', 'STUVW0019S', '105 City Center, Mumbai', false),
('Kyra', 'Chaudhary', 'kyra.c.20@example.com', '919000000020', '100000000020', 'TUVWX0020T', '826 Forest Gate, Noida', false),
('Dev', 'Shah', 'dev.shah.21@domain.in', '919000000021', '100000000021', 'UVWXY0021U', '349 Marina Dr, Mumbai', false),
('Anvi', 'Kapoor', 'anvi.k.22@example.in', '919000000022', '100000000022', 'VWXYZ0022V', '712 Green Valley, Shimla', true),
('Shaurya', 'Bhardwaj', 'shaurya.b.23@mail.in', '919000000023', '100000000023', 'WXYZA0023W', '901 High Tech City, Hyderabad', false),
('Zoya', 'Bhat', 'zoya.b.24@testmail.com', '919000000024', '100000000024', 'XYZAB0024X', '438 Apple Orchard, Srinagar', false),
('Atharv', 'Pillai', 'atharv.p.25@example.com', '919000000025', '100000000025', 'YZABC0025Y', '216 Palm Shore, Thiruvananthapuram', true),
('Sara', 'Menon', 'sara.m.26@domain.com', '919000000026', '100000000026', 'ZABCD0026Z', '573 Backwater View, Alappuzha', false),
('Advait', 'Gill', 'advait.g.27@mail.in', '919000000027', '100000000027', 'ABCDE0027A', '890 Wheat Field, Ludhiana', false),
('Ira', 'Dhillon', 'ira.d.28@example.in', '919000000028', '100000000028', 'BCDEF0028B', '124 Military Cantt, Ambala', true),
('Ayaan', 'Sidhu', 'ayaan.s.29@testmail.in', '919000000029', '100000000029', 'CDEFG0029C', '635 Golden Temple Rd, Amritsar', false),
('Jhiya', 'Sandhu', 'jhiya.s.30@domain.in', '919000000030', '100000000030', 'DEFGH0030D', '902 Rose Garden, Chandigarh', false),
('Aarush', 'Khanna', 'aarush.k.31@example.com', '919000000031', '100000000031', 'EFGHI0031E', '417 Fashion Ave, Mumbai', false),
('Shanaya', 'Suri', 'shanaya.s.32@mail.in', '919000000032', '100000000032', 'FGHIJ0032F', '728 Pearl Tower, Hyderabad', true),
('Kiaan', 'Sarin', 'kiaan.s.33@testmail.com', '919000000033', '100000000033', 'GHIJK0033G', '109 Sun Villa, Jaipur', false),
('Inaya', 'Taneja', 'inaya.t.34@domain.com', '919000000034', '100000000034', 'HIJKL0034H', '842 Moonlight Apts, Delhi', false),
('Ishaan', 'Malik', 'ishaan.m.35@example.in', '919000000035', '100000000035', 'IJKLM0035I', '261 Star Colony, Pune', true),
('Aarya', 'Merchant', 'aarya.m.36@mail.com', '919000000036', '100000000036', 'JKLMN0036J', '584 Blue Lagoon, Goa', false),
('Rehaan', 'Modi', 'rehaan.m.37@testmail.in', '919000000037', '100000000037', 'KLMNO0037K', '395 Business Hub, Ahmedabad', false),
('Vanya', 'Mittal', 'vanya.m.38@example.com', '919000000038', '100000000038', 'LMNOP0038L', '706 Steel City, Jamshedpur', true),
('Dhruv', 'Goel', 'dhruv.g.39@domain.in', '919000000039', '100000000039', 'MNOPQ0039M', '118 IT Park, Bengaluru', false),
('Kavya', 'Bansal', 'kavya.b.40@mail.in', '919000000040', '100000000040', 'NOPQR0040N', '429 Mall Rd, Shimla', false),
('Om', 'Aggarwal', 'om.a.41@testmail.com', '919000000041', '100000000041', 'OPQRS0041O', '742 Hill Top, Mussoorie', false),
('Navya', 'Garg', 'navya.g.42@example.in', '919000000042', '100000000042', 'PQRST0042P', '953 Sunset Blvd, Nagpur', true),
('Yuvraj', 'Gupta', 'yuvraj.g.43@domain.com', '919000000043', '100000000043', 'QRSTU0043Q', '264 River View, Patna', false),
('Amara', 'Bahl', 'amara.b.44@mail.com', '919000000044', '100000000044', 'RSTUV0044R', '575 Palace Rd, Udaipur', false),
('Aditya', 'Sethi', 'aditya.s.45@example.com', '919000000045', '100000000045', 'STUVW0045S', '886 Canal Bank, Kanpur', true),
('Diya', 'Dutta', 'diya.d.46@testmail.in', '919000000046', '100000000046', 'TUVWX0046T', '197 Bridge St, Kolkata', false),
('Ansh', 'Paul', 'ansh.p.47@domain.in', '919000000047', '100000000047', 'UVWXY0047U', '508 North View, Guwahati', false),
('Kiara', 'Roy', 'kiara.r.48@example.in', '919000000048', '100000000048', 'VWXYZ0048V', '819 Forest Rd, Shillong', true),
('Ranbir', 'Sarkar', 'ranbir.s.49@mail.in', '919000000049', '100000000049', 'WXYZA0049W', '130 Tea Garden, Darjeeling', false),
('Riya', 'Banerjee', 'riya.b.50@testmail.com', '919000000050', '100000000050', 'XYZAB0050X', '441 Library Rd, Kolkata', false),
('Siddharth', 'Mukherjee', 'sidd.m.51@example.com', '919000000051', '100000000051', 'YZABC0051Y', '752 Salt Lake, Kolkata', false),
('Prisha', 'Chatterjee', 'prisha.c.52@domain.com', '919000000052', '100000000052', 'ZABCD0052Z', '963 Ballygunge, Kolkata', true),
('Kabir', 'Goswami', 'kabir.g.53@mail.in', '919000000053', '100000000053', 'ABCDE0053A', '274 High Rd, Ranchi', false),
('Sia', 'Bose', 'sia.b.54@example.in', '919000000054', '100000000054', 'BCDEF0054B', '585 Station Rd, Raipur', false),
('Rohan', 'Thakur', 'rohan.t.55@testmail.in', '919000000055', '100000000055', 'CDEFG0055C', '896 Main Bazar, Bhopal', true),
('Ishita', 'Rana', 'ishita.r.56@domain.in', '919000000056', '100000000056', 'DEFGH0056D', '207 Lake View, Udaipur', false),
('Vikram', 'Chauhan', 'vikram.c.57@example.com', '919000000057', '100000000057', 'EFGHI0057E', '518 Fort Rd, Jodhpur', false),
('Meera', 'Rathore', 'meera.r.58@mail.in', '919000000058', '100000000058', 'FGHIJ0058F', '829 Palace View, Jaisalmer', true),
('Abhimanyu', 'Shekhawat', 'abhimanyu.s.59@testmail.com', '919000000059', '100000000059', 'GHIJK0059G', '140 Sand Dune, Bikaner', false),
('Gauri', 'Jadeja', 'gauri.j.60@domain.com', '919000000060', '100000000060', 'HIJKL0060H', '451 Lion Den, Gir', false),
('Tushar', 'Solanki', 'tushar.s.61@example.in', '919000000061', '100000000061', 'IJKLM0061I', '762 Coastal Rd, Dwarka', false),
('Zoya', 'Vaghela', 'zoya.v.62@mail.com', '919000000062', '100000000062', 'JKLMN0062J', '973 Port View, Kandla', true),
('Nikhil', 'Parmar', 'nikhil.p.63@testmail.in', '919000000063', '100000000063', 'KLMNO0063K', '284 Diamond Hub, Surat', false),
('Tara', 'Jhala', 'tara.j.64@example.com', '919000000064', '100000000064', 'LMNOP0064L', '595 Textile Mkt, Surat', false),
('Rishi', 'Gohil', 'rishi.g.65@domain.in', '919000000065', '100000000065', 'MNOPQ0065M', '806 Salt Pan, Bhavnagar', false),
('Anika', 'Zala', 'anika.z.66@mail.in', '919000000066', '100000000066', 'NOPQR0066N', '117 Hill Gate, Junagadh', true),
('Daksh', 'Patil', 'daksh.p.67@testmail.com', '919000000067', '100000000067', 'OPQRS0067O', '428 Orange Grove, Nagpur', false),
('Amrita', 'Deshpande', 'amrita.d.68@example.in', '919000000068', '100000000068', 'PQRST0068P', '739 Cave Rd, Aurangabad', false),
('Samrath', 'Kulkarni', 'samrath.k.69@domain.com', '919000000069', '100000000069', 'QRSTU0069Q', '950 Wine Yard, Nashik', true),
('Leela', 'Joshi', 'leela.j.70@mail.com', '919000000070', '100000000070', 'RSTUV0070R', '261 Temple Rd, Kolhapur', false),
('Madhav', 'Pawar', 'madhav.p.71@example.com', '919000000071', '100000000071', 'STUVW0071S', '572 Sugar Mill, Satara', false),
('Kriti', 'Deshmukh', 'kriti.d.72@testmail.in', '919000000072', '100000000072', 'TUVWX0072T', '883 High Way, Solapur', true),
('Arnav', 'More', 'arnav.m.73@domain.in', '919000000073', '100000000073', 'UVWXY0073U', '194 Creek Side, Thane', false),
('Sana', 'Shinde', 'sana.s.74@example.in', '919000000074', '100000000074', 'VWXYZ0074V', '405 Sea Link, Mumbai', false),
('Yash', 'Gaekwad', 'yash.g.75@mail.in', '919000000075', '100000000075', 'WXYZA0075W', '716 Palace Grounds, Vadodara', false),
('Pihu', 'Chavan', 'pihu.c.76@testmail.com', '919000000076', '100000000076', 'XYZAB0076X', '927 Industrial Estate, Vapi', true),
('Laksh', 'Bhosale', 'laksh.b.77@example.com', '919000000077', '100000000077', 'YZABC0077Y', '238 Green Park, Pune', false),
('Kashvi', 'Sawant', 'kashvi.s.78@domain.com', '919000000078', '100000000078', 'ZABCD0078Z', '549 Hill Crest, Mahabaleshwar', false),
('Parth', 'Rane', 'parth.r.79@mail.in', '919000000079', '100000000079', 'ABCDE0079A', '860 Beach Rd, Ratnagiri', false),
('Mishti', 'Pradhan', 'mishti.p.80@example.in', '919000000080', '100000000080', 'BCDEF0080B', '171 Forest View, Amravati', true),
('Indrajit', 'Mahapatra', 'indra.m.81@testmail.in', '919000000081', '100000000081', 'CDEFG0081C', '482 Temple Rd, Bhubaneswar', false),
('Bani', 'Behera', 'bani.b.82@domain.in', '919000000082', '100000000082', 'DEFGH0082D', '793 Beach St, Puri', false),
('Siddhant', 'Panda', 'sidd.p.83@example.com', '919000000083', '100000000083', 'EFGHI0083E', '104 Coal Mine, Cuttack', false),
('Oishi', 'Dash', 'oishi.d.84@mail.in', '919000000084', '100000000084', 'FGHIJ0084F', '415 Hill Rd, Sambalpur', true),
('Hrithik', 'Sahu', 'hrithik.s.85@testmail.com', '919000000085', '100000000085', 'GHIJK0085G', '726 Lake View, Berhampur', false),
('Aishani', 'Mishra', 'aishani.m.86@domain.com', '919000000086', '100000000086', 'HIJKL0086H', '937 Main Rd, Rourkela', false),
('Ranveer', 'Hegde', 'ranveer.h.87@example.in', '919000000087', '100000000087', 'IJKLM0087I', '248 Coffee Estate, Coorg', true),
('Aditi', 'Shetty', 'aditi.s.88@mail.com', '919000000088', '100000000088', 'JKLMN0088J', '559 Beach Rd, Mangaluru', false),
('Kush', 'Rao', 'kush.r.89@testmail.in', '919000000089', '100000000089', 'KLMNO0089K', '870 Metro St, Bengaluru', false),
('Meher', 'Bhat', 'meher.b.90@example.com', '919000000090', '100000000090', 'LMNOP0090L', '181 Hill View, Hubli', true),
('Nakul', 'Gowda', 'nakul.g.91@domain.in', '919000000091', '100000000091', 'MNOPQ0091M', '492 Farm Land, Mandya', false),
('Sanvi', 'Reddy', 'sanvi.r.92@mail.in', '919000000092', '100000000092', 'NOPQR0092N', '803 IT Park, Gachibowli', false),
('Pranav', 'Naidu', 'pranav.n.93@testmail.com', '919000000093', '100000000093', 'OPQRS0093O', '114 Port Rd, Visakhapatnam', false),
('Anika', 'Choudary', 'anika.c.94@example.in', '919000000094', '100000000094', 'PQRST0094P', '425 Temple Rd, Vijayawada', true),
('Raghav', 'Teja', 'raghav.t.95@domain.com', '919000000095', '100000000095', 'QRSTU0095Q', '736 Canal Rd, Guntur', false),
('Asha', 'Raju', 'asha.r.96@mail.com', '919000000096', '100000000096', 'RSTUV0096R', '947 Beach View, Nellore', false),
('Manish', 'Verma', 'manish.v.97@example.com', '919000000097', '100000000097', 'STUVW0097S', '258 City Center, Warangal', true),
('Juhi', 'Sharma', 'juhi.s.98@testmail.in', '919000000098', '100000000098', 'TUVWX0098T', '569 Lake Rd, Tirupati', false),
('Kartik', 'Agarwal', 'kartik.a.99@domain.in', '919000000099', '100000000099', 'UVWXY0099U', '880 Main St, Kurnool', false),
('Sia', 'Bansal', 'sia.b.100@example.in', '919000000100', '100000000100', 'VWXYZ0100V', '191 High Rd, Nizamabad', true);

-- 7. Insert Initial Donations
INSERT INTO donations (donor_id, amount, currency_code, donation_date, notes) VALUES 
(1, 5000.00, 'INR', CURRENT_TIMESTAMP - INTERVAL '2 days', 'Monthly contribution'),
(2, 100.00, 'INR', CURRENT_TIMESTAMP - INTERVAL '1 day', 'Birthday gift donation'),
(1, 2500.00, 'INR', CURRENT_TIMESTAMP, 'One-time bonus donation'),
(1, 5000.00, 'INR', CURRENT_TIMESTAMP - INTERVAL '2 days', 'Monthly contribution'),
(2, 100.00, 'INR', CURRENT_TIMESTAMP - INTERVAL '1 day', 'Birthday gift donation'),
(1, 2500.00, 'INR', CURRENT_TIMESTAMP, 'One-time bonus donation'),
(1, 5000.00, 'INR', CURRENT_TIMESTAMP - INTERVAL '2 days', 'Monthly contribution'),
(2, 100.00, 'INR', CURRENT_TIMESTAMP - INTERVAL '1 day', 'Birthday gift donation'),
(1, 2500.00, 'INR', CURRENT_TIMESTAMP, 'One-time bonus donation'),
(1, 5000.00, 'INR', CURRENT_TIMESTAMP - INTERVAL '2 days', 'Monthly contribution'),
(2, 100.00, 'INR', CURRENT_TIMESTAMP - INTERVAL '1 day', 'Birthday gift donation'),
(1, 2500.00, 'INR', CURRENT_TIMESTAMP, 'One-time bonus donation'),
(1, 5000.00, 'INR', CURRENT_TIMESTAMP - INTERVAL '2 days', 'Monthly contribution'),
(2, 100.00, 'INR', CURRENT_TIMESTAMP - INTERVAL '1 day', 'Birthday gift donation'),
(1, 2500.00, 'INR', CURRENT_TIMESTAMP, 'One-time bonus donation'),
(1, 5000.00, 'INR', CURRENT_TIMESTAMP - INTERVAL '2 days', 'Monthly contribution'),
(2, 100.00, 'INR', CURRENT_TIMESTAMP - INTERVAL '1 day', 'Birthday gift donation'),
(1, 2500.00, 'INR', CURRENT_TIMESTAMP, 'One-time bonus donation'),
(1, 5000.00, 'INR', CURRENT_TIMESTAMP - INTERVAL '2 days', 'Monthly contribution'),
(2, 100.00, 'INR', CURRENT_TIMESTAMP - INTERVAL '1 day', 'Birthday gift donation'),
(1, 2500.00, 'INR', CURRENT_TIMESTAMP, 'One-time bonus donation'),
(1, 5000.00, 'INR', CURRENT_TIMESTAMP - INTERVAL '2 days', 'Monthly contribution'),
(2, 100.00, 'INR', CURRENT_TIMESTAMP - INTERVAL '1 day', 'Birthday gift donation'),
(1, 2500.00, 'INR', CURRENT_TIMESTAMP, 'One-time bonus donation'),
(1, 5000.00, 'INR', CURRENT_TIMESTAMP - INTERVAL '2 days', 'Monthly contribution'),
(2, 100.00, 'INR', CURRENT_TIMESTAMP - INTERVAL '1 day', 'Birthday gift donation'),
(1, 2500.00, 'INR', CURRENT_TIMESTAMP, 'One-time bonus donation'),
(1, 5000.00, 'INR', CURRENT_TIMESTAMP - INTERVAL '2 days', 'Monthly contribution'),
(2, 100.00, 'INR', CURRENT_TIMESTAMP - INTERVAL '1 day', 'Birthday gift donation'),
(1, 2500.00, 'INR', CURRENT_TIMESTAMP, 'One-time bonus donation'),
(1, 5000.00, 'INR', CURRENT_TIMESTAMP - INTERVAL '2 days', 'Monthly contribution'),
(2, 100.00, 'INR', CURRENT_TIMESTAMP - INTERVAL '1 day', 'Birthday gift donation'),
(1, 2500.00, 'INR', CURRENT_TIMESTAMP, 'One-time bonus donation');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM donations;
DELETE FROM members;
DELETE FROM currencies;
DELETE FROM user_roles;
DELETE FROM users;
DELETE FROM roles;
-- +goose StatementEnd