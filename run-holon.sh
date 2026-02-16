#!/bin/bash
# Holon 演示脚本 - 在 Docker 容器中运行

# 方式 1: 使用 Docker 直接运行
# 需要设置环境变量:
#   ANTHROPIC_AUTH_TOKEN - Anthropic API Token

echo "=== Holon 演示 ==="
echo ""
echo "当前项目状态:"
echo "- main.go: Calculator 实现，Multiply 方法有 bug"
echo "- main_test.go: 测试用例，TestMultiply 会失败"
echo ""

# 显示当前测试结果
echo ">>> 当前测试结果:"
cd /workspace
go test -v

echo ""
echo ">>> 使用 Holon 修复 bug..."
echo "命令: holon run --spec holon-spec.yaml --image golang:1.22 --workspace ."
echo ""

# Holon 会:
# 1. 读取 holon-spec.yaml 了解任务
# 2. 在 Docker 容器中启动 Claude Code
# 3. AI 分析代码，找到 bug
# 4. 修复 Multiply 函数
# 5. 运行测试验证
# 6. 产出 diff.patch 和 summary.md

# holon run \
#   --spec holon-spec.yaml \
#   --image golang:1.22 \
#   --workspace . \
#   --output ./holon-output

echo ">>> Holon 产出物:"
echo "- holon-output/diff.patch   : 代码变更补丁"
echo "- holon-output/summary.md   : AI 执行摘要"
echo "- holon-output/manifest.json: 运行元数据"
