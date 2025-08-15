import requests
import json

# API基础URL
BASE_URL = "http://localhost:8080/api"

# 测试用户数据
TEST_USER = {
    "name": "testuser",
    "password": "testpassword123"
}

def register_user():
    """测试用户注册"""
    print("=== 测试用户注册 ===")
    try:
        response = requests.post(
            f"{BASE_URL}/register",
            json=TEST_USER
        )
        print(f"状态码: {response.status_code}")
        print(f"响应: {response.json()}")
        return response.status_code == 200
    except Exception as e:
        print(f"注册请求失败: {e}")
        return False

def login_user():
    """测试用户登录"""
    print("\n=== 测试用户登录 ===")
    try:
        response = requests.post(
            f"{BASE_URL}/login",
            json=TEST_USER
        )
        print(f"状态码: {response.status_code}")
        print(f"响应: {response.json()}")
        
        if response.status_code == 200:
            token = response.json().get('data', {}).get('token')
            if token:
                print(f"获取到JWT Token: {token[:20]}...")
                return token
            else:
                print("响应中未找到token")
                return None
        return None
    except Exception as e:
        print(f"登录请求失败: {e}")
        return None

def get_user_by_id(token, user_id=1):
    """测试需要认证的API端点"""
    print("\n=== 测试需要认证的API端点 ===")
    try:
        headers = {
            "Authorization": f"Bearer {token}"
        }
        response = requests.get(
            f"{BASE_URL}/user/id?id={user_id}",
            headers=headers
        )
        print(f"状态码: {response.status_code}")
        print(f"响应: {response.json()}")
        return response.status_code == 200
    except Exception as e:
        print(f"获取用户信息失败: {e}")
        return False

def get_user_by_id_without_auth():
    """测试无认证访问受保护的API端点"""
    print("\n=== 测试无认证访问受保护的API端点 ===")
    try:
        response = requests.get(
            f"{BASE_URL}/user/id?id=1"
        )
        print(f"状态码: {response.status_code}")
        print(f"响应: {response.json()}")
        return response.status_code == 401
    except Exception as e:
        print(f"无认证访问测试失败: {e}")
        return False

def get_user_by_name():
    """测试公开的API端点"""
    print("\n=== 测试公开的API端点 ===")
    try:
        response = requests.get(
            f"{BASE_URL}/user/name?name={TEST_USER['name']}"
        )
        print(f"状态码: {response.status_code}")
        print(f"响应: {response.json()}")
        return response.status_code in [200, 404]  # 200表示找到用户，404表示未找到
    except Exception as e:
        print(f"获取用户信息失败: {e}")
        return False

def main():
    print("开始测试认证模块...")
    
    # 1. 测试注册
    if not register_user():
        print("注册失败，测试终止")
        return
    
    # 2. 测试登录
    token = login_user()
    if not token:
        print("登录失败，测试终止")
        return
    
    # 3. 测试需要认证的API
    if not get_user_by_id(token):
        print("认证API测试失败")
    
    # 4. 测试无认证访问受保护的API
    if not get_user_by_id_without_auth():
        print("无认证访问测试失败")
    
    # 5. 测试公开API
    if not get_user_by_name():
        print("公开API测试失败")
    
    print("\n所有测试完成!")

if __name__ == "__main__":
    main()