//
//  ContentView.swift
//  finalProjectV1
//
//  Created by Lwj on 2021/11/21.
//

import SwiftUI

struct ContentView: View {
    
    @State private var name = ""
    @State private var pswd = ""
    @State private var flag: Int? = nil
    @State private var authSuccess: Bool = false
    @State private var authfailed: Bool = false
    
    
    @StateObject var viewModel = LoginViewModelImpl(service: LoginServiceImpl())
    
    var body: some View {
        NavigationView(content: {
            VStack{
                
                LoginTextContent()
                
                UserImage()
                
                UserNameTextField(name: $name)
                PasswordSecureField(pswd: $pswd)
                
                if self.authfailed {
                    Text("failed !  please retry")
                        .font(.body)
                        .foregroundColor(.red)
                        .padding(.bottom, 20)
                }
                
                NavigationLink(destination: HomeView(),tag: 1,selection: $flag){
                    Button(action: {
                        if viewModel.checkUserNameAndPassword(name, pswd) {
                            self.flag = 1
                        } else {
                            self.authfailed = true
                        }
                    }, label: {
                        LoginButtonContent()
                    })
                }
                
                              
            }
            .padding()
        })
    }
}

struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
    }
}

struct LoginTextContent: View {
    var body: some View {
        Text("校园App 登陆界面")
            .fontWeight(/*@START_MENU_TOKEN@*/.bold/*@END_MENU_TOKEN@*/)
            .padding(.bottom, 30)
            .font(.largeTitle)
    }
}

struct UserImage: View {
    var body: some View {
        Image("UserImage")
            .resizable()
            .aspectRatio(contentMode: .fill)
            .frame(width: 150, height: 150)
            .clipped()
            .cornerRadius(150)
            .padding(.bottom, 50)
    }
}

struct UserNameTextField: View {
    @Binding var name: String
    var body: some View {
        TextField("name:", text: $name)
            .padding(.all)
            .background(.thinMaterial)
            .padding(.bottom, 20)
    }
}

struct PasswordSecureField: View {
    @Binding var pswd: String
    var body: some View {
        SecureField("password:", text: $pswd)
            .padding(/*@START_MENU_TOKEN@*/.all/*@END_MENU_TOKEN@*/)
            .background(.thinMaterial)
            .padding(.bottom, 20)
    }
}

struct LoginButtonContent: View {
    var body: some View {
        Text("Login")
            .font(.headline)
            .foregroundColor(.white)
            .padding()
            .frame(width: 220, height: 60)
            .background(.black)
            .cornerRadius(35)
    }
}



