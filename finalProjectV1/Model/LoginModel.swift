//
//  LoginModel.swift
//  finalProjectV1
//
//  Created by Lwj on 2021/11/21.
//

import Foundation

struct jsonPostSuccess: Codable {
    
    var Message: String
    var Status: Int
    var token: String
//        "Message": "密码正确",
//        "Status": 0,
//        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiMTIzIiwicGhvbmUiOiIzIiwiZXhwIjoxNjM3NzIyNjgzLCJpc3MiOiJhZG1pbi1zZXJ2ZXIifQ.Kq_CM6tNsJuHUwRA4p5_I_of4iQ0p0k4ZpaNvYWbOaY"
//
}

struct jsonPostFailed: Codable {
    
    var Message: String
    var Status: Int
//    var token: String
//        "Message": "密码正确",
//        "Status": 0,
//        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiMTIzIiwicGhvbmUiOiIzIiwiZXhwIjoxNjM3NzIyNjgzLCJpc3MiOiJhZG1pbi1zZXJ2ZXIifQ.Kq_CM6tNsJuHUwRA4p5_I_of4iQ0p0k4ZpaNvYWbOaY"
//
}

class CheckLogin {
    
//    func postDataLog() {
//        guard let url = URL(string: "http://hutaowlp.xyz:114/api/login") else { return }
//        let title: String = "123"
//        let bar: String = "123"
//        //let userId = 1
//        let body: [String: Any] = ["Username": title, "Password": bar]
//        let finalData = try! JSONSerialization.data(withJSONObject: body)
//        var request = URLRequest(url: url)
//        request.httpBody = finalData
//        request.httpMethod = "POST"
//        request.setValue("application/json", forHTTPHeaderField: "Content-Type")
//
//        URLSession.shared.dataTask(with: request) { (data, res, err) in
//            do {
//                if let data = data {
//                    print(data)
//                    print("--------------------")
//                    let resultS = try JSONDecoder().decode(jsonPostSuccess.self, from: data)
//
//                    print(resultS)
//
//                } else {
//                    print("No data")
//                }
//            } catch (let error) {
//                print(error.localizedDescription)
//            }
//        }.resume()
//    }
    
    
    
    func checkUserNameAndPassword(_ name: String, _ pswd: String, completion: @escaping (Int) -> Bool ) {
        
        guard let url = URL(string: "http://hutaowlp.xyz:114/api/login") else { return }
        let title: String = name
        let bar: String = pswd
        //let userId = 1
        let body: [String: Any] = ["Username": title, "Password": bar]
        let finalData = try! JSONSerialization.data(withJSONObject: body, options: .fragmentsAllowed)
        var request = URLRequest(url: url)
        request.httpBody = finalData
        request.httpMethod = "POST"
        request.setValue("application/json", forHTTPHeaderField: "Content-Type")
        
        URLSession.shared.dataTask(with: request) { data, res, err in
            do {
                if let data = data {
                    print(data)
                    print("--------------------")
                    let resultS = try JSONDecoder().decode(jsonPostSuccess.self, from: data)
                    
//                    DispatchQueue.main.async {
                    completion(resultS.Status)
//                    }
                    
                } else {
                    print("No data")
                }
            } catch (let error) {
                print(error.localizedDescription)
            }
        }.resume()
    }
    
}
