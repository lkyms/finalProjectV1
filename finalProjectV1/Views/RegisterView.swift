//
//  HomeView.swift
//  finalProjectV1
//
//  Created by Lwj on 2021/11/21.
//

import SwiftUI

struct RegisterView: View {
    
    @State private var name: String = ""
    @State private var pswd: String = ""
    
    @State var jobIndex = 0
    var Jobs = ["Student", "Teacher", "Schoolmate"]
    
    var body: some View {
        VStack {
            RegisterTextContent()
            RegisterImage()
                RegisterUserNameTextField(name: $name)
                RegisterPasswordSecureField(pswd: $pswd)
            HStack{
                Text("what are you ?")
                    .padding()
                Spacer()
                Picker("what are you ?", selection: self.$jobIndex) {
                    ForEach(self.Jobs.indices) { num in
                        Text(self.Jobs[num])
                            .font(.headline)
                            .foregroundColor(.black)
                    }
                }.padding(.horizontal, 50)
            }
        }
    }
}

struct RegisterView_Previews: PreviewProvider {
    static var previews: some View {
        RegisterView()
    }
}
struct RegisterTextContent: View {
    var body: some View {
        Text("校园App 注册界面")
            .fontWeight(/*@START_MENU_TOKEN@*/.bold/*@END_MENU_TOKEN@*/)
            .padding(.bottom, 30)
            .font(.largeTitle)
    }
}

struct RegisterImage: View {
    var body: some View {
        Image("RegisterImage")
            .resizable()
            .aspectRatio(contentMode: .fill)
            .frame(width: 150, height: 150)
            .clipped()
            .cornerRadius(150)
            .padding(.bottom, 50)
    }
}

struct RegisterUserNameTextField: View {
    @Binding var name: String
    var body: some View {
        TextField("name:", text: $name)
            .padding(.all)
            .background(.thinMaterial)
            .padding(.bottom, 20)
    }
}

struct RegisterPasswordSecureField: View {
    @Binding var pswd: String
    var body: some View {
        SecureField("password:", text: $pswd)
            .padding(/*@START_MENU_TOKEN@*/.all/*@END_MENU_TOKEN@*/)
            .background(.thinMaterial)
            .padding(.bottom, 20)
    }
}
